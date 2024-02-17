package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type JavaNetUrlConnectionObjectWrapper struct {
	*jvm.ObjectWrapper
}

func (j *JavaNetUrlConnectionObjectWrapper) AddRequestProperty(key string, value string) mo.Result[struct{}] {
	r, err := j.CallVoidA("addRequestProperty", key, value).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) Connect() mo.Result[struct{}] {
	r, err := j.CallVoidA("connect").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetConnectTimeout() mo.Result[int] {
	r, err := j.CallIntA("getConnectTimeout").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetContentEncoding() mo.Result[string] {
	r, err := j.CallStringA("getContentEncoding").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetContentLength() mo.Result[int] {
	r, err := j.CallIntA("getContentLength").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetContentLengthLong() mo.Result[int64] {
	r, err := j.CallLongA("getContentLengthLong").Get()
	if err != nil {
		return mo.Err[int64](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetContentType() mo.Result[string] {
	r, err := j.CallStringA("getContentType").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetDate() mo.Result[int64] {
	r, err := j.CallLongA("getDate").Get()
	if err != nil {
		return mo.Err[int64](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetDefaultUseCaches() mo.Result[bool] {
	r, err := j.CallBoolA("getDefaultUseCaches").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetDoInput() mo.Result[bool] {
	r, err := j.CallBoolA("getDoInput").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetDoOutput() mo.Result[bool] {
	r, err := j.CallBoolA("getDoOutput").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetExpiration() mo.Result[int64] {
	r, err := j.CallLongA("getExpiration").Get()
	if err != nil {
		return mo.Err[int64](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetHeaderField_n(n int) mo.Result[string] {
	r, err := j.CallStringA("getHeaderField", n).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetHeaderField_name(name string) mo.Result[string] {
	r, err := j.CallStringA("getHeaderField", name).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetHeaderFieldKey_n(n int) mo.Result[string] {
	r, err := j.CallStringA("getHeaderFieldKey", n).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetHeaderFields() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("getHeaderFields").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetInputStream() mo.Result[*JavaIoInputStreamObjectWrapper] {
	r, err := j.CallObjectA("getInputStream").Get()
	if err != nil {
		return mo.Err[*JavaIoInputStreamObjectWrapper](err)
	}
	return mo.Ok(JavaIoInputStreamWithJniPtr(r.JniPtr()))
}
func (j *JavaNetUrlConnectionObjectWrapper) GetOutputStream() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("getOutputStream").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetReadTimeout() mo.Result[int] {
	r, err := j.CallIntA("getReadTimeout").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetRequestProperties() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("getRequestProperties").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetRequestProperty_key(key string) mo.Result[string] {
	r, err := j.CallStringA("getRequestProperty", key).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetURL() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("getURL").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) GetUseCaches() mo.Result[bool] {
	r, err := j.CallBoolA("getUseCaches").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(r)
}

func (j *JavaNetUrlConnectionObjectWrapper) SetConnectTimeout(timeout int) mo.Result[struct{}] {
	r, err := j.CallVoidA("setConnectTimeout", timeout).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) SetDoInput(doinput bool) mo.Result[struct{}] {
	r, err := j.CallVoidA("setDoInput", doinput).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) SetDoOutput(dooutput bool) mo.Result[struct{}] {
	r, err := j.CallVoidA("setDoOutput", dooutput).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) SetReadTimeout(timeout int) mo.Result[struct{}] {
	r, err := j.CallVoidA("setReadTimeout", timeout).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) SetRequestProperty(key string, value string) mo.Result[struct{}] {
	r, err := j.CallVoidA("setRequestProperty", key, value).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaNetUrlConnectionObjectWrapper) SetUseCaches(usecaches bool) mo.Result[struct{}] {
	r, err := j.CallVoidA("setUseCaches", usecaches).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}

func JavaNetUrlConnectionWithJniPtr(ptr uintptr) *JavaNetUrlConnectionObjectWrapper {
	return &JavaNetUrlConnectionObjectWrapper{
		ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr),
	}
}
