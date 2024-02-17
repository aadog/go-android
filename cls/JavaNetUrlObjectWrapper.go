package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type JavaNetUrlObjectWrapper struct {
	*jvm.ObjectWrapper
}

func (j *JavaNetUrlObjectWrapper) GetDefaultPort() mo.Result[int] {
	s, err := j.CallIntA("getDefaultPort").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(s)
}

func (j *JavaNetUrlObjectWrapper) GetAuthority() mo.Result[string] {
	s, err := j.CallStringA("getAuthority").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) GetFile() mo.Result[string] {
	s, err := j.CallStringA("getFile").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) GetHost() mo.Result[string] {
	s, err := j.CallStringA("getHost").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) GetPath() mo.Result[string] {
	s, err := j.CallStringA("getPath").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) GetProtocol() mo.Result[string] {
	s, err := j.CallStringA("GetProtocol").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) GetQuery() mo.Result[string] {
	s, err := j.CallStringA("GetQuery").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) GetUserInfo() mo.Result[string] {
	s, err := j.CallStringA("getUserInfo").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) OpenStream() mo.Result[*jvm.ObjectWrapper] {
	s, err := j.CallObjectA("openStream").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(s)
}
func (j *JavaNetUrlObjectWrapper) OpenConnection() mo.Result[*JavaNetUrlConnectionObjectWrapper] {
	s, err := j.CallObjectA("openConnection").Get()
	if err != nil {
		return mo.Err[*JavaNetUrlConnectionObjectWrapper](err)
	}
	return mo.Ok(JavaNetUrlConnectionWithJniPtr(s.JniPtr()))
}
func (j *JavaNetUrlObjectWrapper) OpenConnection_proxy(proxy *jvm.ObjectWrapper) mo.Result[*JavaNetUrlConnectionObjectWrapper] {
	s, err := j.CallObjectA("openConnection", proxy).Get()
	if err != nil {
		return mo.Err[*JavaNetUrlConnectionObjectWrapper](err)
	}
	return mo.Ok(JavaNetUrlConnectionWithJniPtr(s.JniPtr()))
}
func (j *JavaNetUrlObjectWrapper) OpenHttpConnection() mo.Result[*ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper] {
	s, err := j.CallObjectA("openConnection").Get()
	if err != nil {
		return mo.Err[*ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper](err)
	}
	return mo.Ok(ComAndroidOkhttpInternalHucHttpURLConnectionImplWithJniPtr(s.JniPtr()))
}
func (j *JavaNetUrlObjectWrapper) OpenHttpsConnection() mo.Result[*ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper] {
	s, err := j.CallObjectA("openConnection").Get()
	if err != nil {
		return mo.Err[*ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper](err)
	}
	return mo.Ok(ComAndroidOkhttpInternalHucHttpsURLConnectionImplWithJniPtr(s.JniPtr()))
}
func (j *JavaNetUrlObjectWrapper) OpenHttpConnection_proxy(proxy *jvm.ObjectWrapper) mo.Result[*ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper] {
	s, err := j.CallObjectA("openConnection", proxy).Get()
	if err != nil {
		return mo.Err[*ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper](err)
	}
	return mo.Ok(ComAndroidOkhttpInternalHucHttpURLConnectionImplWithJniPtr(s.JniPtr()))
}
func (j *JavaNetUrlObjectWrapper) OpenHttpsConnection_proxy(proxy *jvm.ObjectWrapper) mo.Result[*ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper] {
	s, err := j.CallObjectA("openConnection", proxy).Get()
	if err != nil {
		return mo.Err[*ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper](err)
	}
	return mo.Ok(ComAndroidOkhttpInternalHucHttpsURLConnectionImplWithJniPtr(s.JniPtr()))
}

func JavaNetUrlObjectWrapperWithJniPtr(ptr uintptr) *JavaNetUrlObjectWrapper {
	return &JavaNetUrlObjectWrapper{
		ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr),
	}
}
