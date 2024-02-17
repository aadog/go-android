package cls

import (
	"github.com/aadog/go-ndk/jvm"
	"github.com/samber/mo"
)

type JavaIoBufferedReaderObjectWrapper struct {
	*jvm.ObjectWrapper
}

func (j *JavaIoBufferedReaderObjectWrapper) Close() mo.Result[struct{}] {
	r, err := j.CallVoidA("close").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaIoBufferedReaderObjectWrapper) Lines() mo.Result[*jvm.ObjectWrapper] {
	r, err := j.CallObjectA("lines").Get()
	if err != nil {
		return mo.Err[*jvm.ObjectWrapper](err)
	}
	return mo.Ok(r)
}
func (j *JavaIoBufferedReaderObjectWrapper) Mark(readAheadLimit int) mo.Result[struct{}] {
	r, err := j.CallVoidA("mark", readAheadLimit).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}

func (j *JavaIoBufferedReaderObjectWrapper) MarkSupported() mo.Result[bool] {
	n, err := j.CallBoolA("markSupported").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(n)
}

func (j *JavaIoBufferedReaderObjectWrapper) Read() mo.Result[int] {
	n, err := j.CallIntA("read").Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}

func (j *JavaIoBufferedReaderObjectWrapper) Read_cbuf_off_len(cbuf *jvm.ObjectWrapper, off int, _len int) mo.Result[int] {
	n, err := j.CallIntA("read", cbuf, off, _len).Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(n)
}

func (j *JavaIoBufferedReaderObjectWrapper) ReadLine() mo.Result[*string] {
	n, err := j.CallPStringA("readLine").Get()
	if err != nil {
		return mo.Err[*string](err)
	}
	return mo.Ok(n)
}

func (j *JavaIoBufferedReaderObjectWrapper) Ready() mo.Result[bool] {
	n, err := j.CallBoolA("ready").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(n)
}
func (j *JavaIoBufferedReaderObjectWrapper) Reset() mo.Result[struct{}] {
	r, err := j.CallVoidA("reset").Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(r)
}
func (j *JavaIoBufferedReaderObjectWrapper) Skip(n int64) mo.Result[int64] {
	n, err := j.CallLongA("skip", n).Get()
	if err != nil {
		return mo.Err[int64](err)
	}
	return mo.Ok(n)
}

func JavaIoBufferedReaderWithJniPtr(ptr uintptr) *JavaIoBufferedReaderObjectWrapper {
	return &JavaIoBufferedReaderObjectWrapper{
		ObjectWrapper: jvm.ObjectWrapperWithJniPtr(ptr),
	}
}
