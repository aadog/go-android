package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type JavaIoInputStreamObjectWrapper struct {
	*jvm.ObjectWrapper
}

func (j *JavaIoInputStreamObjectWrapper) Available() mo.Result[int] {
	n, err := j.CallIntA("available").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamObjectWrapper) Close() mo.Result[struct{}] {
	r, err := j.CallVoidA("close").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaIoInputStreamObjectWrapper) Mark(readlimit int) mo.Result[struct{}] {
	r, err := j.CallVoidA("mark", readlimit).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaIoInputStreamObjectWrapper) MarkSupported() mo.Result[bool] {
	n, err := j.CallBoolA("markSupported").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamObjectWrapper) Read() mo.Result[int] {
	n, err := j.CallIntA("read").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamObjectWrapper) Read_b(b *jvm.ObjectWrapper) mo.Result[int] {
	n, err := j.CallIntA("read", b).Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamObjectWrapper) Read_b_off_len(b *jvm.ObjectWrapper, off int, _len int) mo.Result[int] {
	n, err := j.CallIntA("read", b, off, _len).Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoInputStreamObjectWrapper) Reset() mo.Result[struct{}] {
	r, err := j.CallVoidA("reset").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaIoInputStreamObjectWrapper) Skip(n int64) mo.Result[int64] {
	n, err := j.CallLongA("skip", n).Get()
	if err != nil {
		return mo.Err[int64](err)
	}
	return mo.Ok(n)
}

func JavaIoInputStreamWithJniPtr(ptr uintptr) *JavaIoInputStreamObjectWrapper {
	return &JavaIoInputStreamObjectWrapper{
		ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr),
	}
}
