package jvm

import (
	"errors"
	"github.com/aadog/go-ndk/jni"
	"github.com/samber/mo"
)

type JniCallRet interface {
	*ObjectWrapper | string | struct{} | int | *string | int64 | bool
}

func GetClassName(ptr jni.Jclass) mo.Result[string] {
	env := LocalThreadJavaEnv()
	objCls, err := env.GetObjectClass(ptr).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(objCls)
	getNameMethodId, err := env.GetMethodID(objCls, "getName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	jstringName, err := env.CallObjectMethodA(ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}

func ObjectWrapperCall[T JniCallRet](o *ObjectWrapper, funcName string, args ...any) mo.Result[any] {
	env := LocalThreadJavaEnv()
	objCls, err := o.Class().Get()
	if err != nil {
		return mo.Err[any](err)
	}
	method, err := objCls.EnumMatchMethod(funcName, SumGoArgsType(args...)...).Get()
	if err != nil {
		return mo.Err[any](err)
	}
	defer method.Free()

	methodId, err := env.FromReflectedMethod(method.JniPtr()).Get()
	if err != nil {
		return mo.Err[any](err)
	}
	jArgs := make([]jni.Jvalue, 0)
	for _, arg := range args {
		jval, needDelete := ConvertAnyArgToJValueArg(arg)
		if needDelete {
			defer env.DeleteLocalRef(jni.Jobject(jval))
		}
		jArgs = append(jArgs, jval)
	}
	var inputType T

	switch any(inputType).(type) {
	case *ObjectWrapper:
		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		objWrapper := ObjectWrapperWithJniPtr(obj)
		return mo.Ok(any(objWrapper))
	case *ClassWrapper:
		panic(errors.New("system error"))
		//obj, err := o.CallObjectMethodA(methodId, jArgs...).Get()
		//if err != nil {
		//	return mo.Err[any](err)
		//}
		//defer env.DeleteLocalRef(obj)
		//cls := o.Class()
		//return mo.Ok(any(cls))
	case struct{}:
		_, err := env.CallVoidMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(struct{}{}))
	case string:
		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			return mo.Ok(any(""))
		}
		defer env.DeleteLocalRef(obj)
		jstring := env.GetStringUTF(obj)
		return mo.Ok(any(jstring))
	case *string:
		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			var p *string = nil
			return mo.Ok(any(p))
		}
		defer env.DeleteLocalRef(obj)
		jstring := string(env.GetStringUTF(obj))
		return mo.Ok(any(&jstring))
	case int:
		n, err := env.CallIntMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case int64:
		n, err := env.CallIntMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case bool:
		n, err := env.CallBooleanMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	default:
		panic(errors.New("暂不支持"))
	}
}
func GetPrimitiveClass(className string) mo.Result[jni.Jclass] {
	env := LocalThreadJavaEnv()
	classClass, err := env.FindClass("java.lang.Class").Get()
	if err != nil {
		return mo.Errf[jni.Jclass]("find class error:%v", err)
	}
	defer env.DeleteLocalRef(classClass)
	methodId, err := env.GetStaticMethodID(classClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;").Get()
	if err != nil {
		return mo.Errf[jni.Jclass]("find class error:%v", err)
	}
	ss := env.NewString(className)
	defer env.DeleteLocalRef(ss)
	cls, err := env.CallStaticObjectMethodA(classClass, methodId, jni.Jvalue(ss)).Get()
	if err != nil {
		return mo.Errf[jni.Jclass]("find class error:%v", err)
	}
	return mo.Ok(cls)
}
