package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type JavaIoInputStreamReaderObjectWrapper struct {
	*jvm.ObjectWrapper
}

func (j *JavaIoInputStreamReaderObjectWrapper) Close() mo.Result[struct{}] {
	r, err := j.CallVoidA("close").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}

func (j *JavaIoInputStreamReaderObjectWrapper) GetEncoding() mo.Result[string] {
	n, err := j.CallStringA("getEncoding").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamReaderObjectWrapper) Read() mo.Result[int] {
	n, err := j.CallIntA("read").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamReaderObjectWrapper) Read_cbuf_off_len(cbuf *jvm.ObjectWrapper, off int, _len int) mo.Result[int] {
	n, err := j.CallIntA("read", cbuf, off, _len).Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamReaderObjectWrapper) Ready() mo.Result[bool] {
	n, err := j.CallBoolA("ready").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(n)
}

func JavaIoInputStreamReaderWithJniPtr(ptr uintptr) *JavaIoInputStreamReaderObjectWrapper {
	return &JavaIoInputStreamReaderObjectWrapper{
		ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr),
	}
}
