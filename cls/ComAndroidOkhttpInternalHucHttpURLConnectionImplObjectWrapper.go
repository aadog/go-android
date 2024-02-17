package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper struct {
	*JavaNetUrlConnectionObjectWrapper
}

func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) Disconnect() mo.Result[struct{}] {
	r, err := j.CallVoidA("disconnect").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) GetRequestMethod() mo.Result[string] {
	r, err := j.CallStringA("getRequestMethod").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) GetResponseCode() mo.Result[int] {
	r, err := j.CallIntA("getResponseCode").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) GetResponseMessage() mo.Result[string] {
	r, err := j.CallStringA("getResponseMessage").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) SetChunkedStreamingMode(chunklen int) mo.Result[struct{}] {
	r, err := j.CallVoidA("setChunkedStreamingMode", chunklen).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) SetFixedLengthStreamingMode(contentLength int64) mo.Result[struct{}] {
	r, err := j.CallVoidA("setFixedLengthStreamingMode", contentLength).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) SetInstanceFollowRedirects(followRedirects bool) mo.Result[struct{}] {
	r, err := j.CallVoidA("setInstanceFollowRedirects", followRedirects).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) SetRequestMethod(method string) mo.Result[struct{}] {
	r, err := j.CallVoidA("setRequestMethod", method).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper) UsingProxy() mo.Result[bool] {
	r, err := j.CallBoolA("usingProxy").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(r)
}

func ComAndroidOkhttpInternalHucHttpURLConnectionImplWithJniPtr(ptr uintptr) *ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper {
	return &ComAndroidOkhttpInternalHucHttpURLConnectionImplObjectWrapper{
		JavaNetUrlConnectionObjectWrapper: &JavaNetUrlConnectionObjectWrapper{ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr)},
	}
}
