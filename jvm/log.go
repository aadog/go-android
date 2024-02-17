package jvm

/*
#include <android/log.h>
*/
import "C"
import (
	"fmt"
	"github.com/aadog/go-ndk/utils"
)

func LogError(tag, format string, v ...interface{}) {
	ctag := (*C.char)(utils.CString(tag))
	cstr := (*C.char)(utils.CString(fmt.Sprintf(format, v...)))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
}

func LogInfo(tag, format string, v ...interface{}) {
	ctag := (*C.char)(utils.CString(tag))
	cstr := (*C.char)(utils.CString(fmt.Sprintf(format, v...)))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
}
