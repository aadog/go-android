package android

/*
#include <android/log.h>
*/
import "C"
import (
	"fmt"
	"github.com/aadog/go-android/utils"
)

func AndroidLogError(tag, format string, v ...interface{}) {
	ctag := utils.CString(tag)
	cstr := utils.CString(fmt.Sprintf(format, v...))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
}

func AndroidLogInfo(tag, format string, v ...interface{}) {
	ctag := utils.CString(tag)
	cstr := utils.CString(fmt.Sprintf(format, v...))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
}
