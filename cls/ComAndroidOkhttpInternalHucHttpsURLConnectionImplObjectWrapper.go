package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper struct {
	*JavaNetUrlConnectionObjectWrapper
}

func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) Disconnect() mo.Result[struct{}] {
	r, err := j.CallVoidA("disconnect").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) GetRequestMethod() mo.Result[string] {
	r, err := j.CallStringA("getRequestMethod").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) GetResponseCode() mo.Result[int] {
	r, err := j.CallIntA("getResponseCode").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) GetResponseMessage() mo.Result[string] {
	r, err := j.CallStringA("getResponseMessage").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) SetChunkedStreamingMode(chunklen int) mo.Result[struct{}] {
	r, err := j.CallVoidA("setChunkedStreamingMode", chunklen).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) SetFixedLengthStreamingMode(contentLength int64) mo.Result[struct{}] {
	r, err := j.CallVoidA("setFixedLengthStreamingMode", contentLength).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) SetInstanceFollowRedirects(followRedirects bool) mo.Result[struct{}] {
	r, err := j.CallVoidA("setInstanceFollowRedirects", followRedirects).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) SetRequestMethod(method string) mo.Result[struct{}] {
	r, err := j.CallVoidA("setRequestMethod", method).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) UsingProxy() mo.Result[bool] {
	r, err := j.CallBoolA("usingProxy").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) GetCipherSuite() mo.Result[string] {
	r, err := j.CallStringA("getCipherSuite").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) SetHostnameVerifier(hostnameVerifier *jvm.ObjectWrapper) mo.Result[struct{}] {
	r, err := j.CallVoidA("setHostnameVerifier", hostnameVerifier).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) GetHostnameVerifier() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("getHostnameVerifier").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) SetSSLSocketFactory(sslSocketFactory *jvm.ObjectWrapper) mo.Result[struct{}] {
	r, err := j.CallVoidA("setSSLSocketFactory", sslSocketFactory).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper) GetSSLSocketFactory() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("getSSLSocketFactory").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}

func ComAndroidOkhttpInternalHucHttpsURLConnectionImplWithJniPtr(ptr uintptr) *ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper {
	return &ComAndroidOkhttpInternalHucHttpsURLConnectionImplObjectWrapper{
		JavaNetUrlConnectionObjectWrapper: &JavaNetUrlConnectionObjectWrapper{ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr)},
	}
}
