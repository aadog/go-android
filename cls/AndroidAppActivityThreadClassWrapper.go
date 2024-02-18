package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type AndroidAppActivityThreadClassWrapper struct {
	*jvm.ClassWrapper
}

func (j *AndroidAppActivityThreadClassWrapper) CurrentActivityThread() mo.Result[*jvm.ObjectWrapper] {
	c, err := j.CallStaticObjectA("currentActivityThread").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(jvm.ObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *AndroidAppActivityThreadClassWrapper) GetApplication() mo.Result[*jvm.ObjectWrapper] {
	c, err := j.CallStaticObjectA("getApplication").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(jvm.ObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *AndroidAppActivityThreadClassWrapper) GetPackageManager() mo.Result[*jvm.ObjectWrapper] {
	c, err := j.CallStaticObjectA("getPackageManager").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(jvm.ObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *AndroidAppActivityThreadClassWrapper) GetIntentBeingBroadcast() mo.Result[*jvm.ObjectWrapper] {
	c, err := j.CallStaticObjectA("getIntentBeingBroadcast").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(jvm.ObjectWrapperWithJniPtr(c.JniPtr()))
}

func AndroidAppActivityThreadClass() *AndroidAppActivityThreadClassWrapper {
	return &AndroidAppActivityThreadClassWrapper{
		ClassWrapper: jvm.Use("android.app.ActivityThread").MustGet(),
	}
}
