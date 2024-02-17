package dlfcn

/*
#include <dlfcn.h>
*/
import "C"
import (
	"unsafe"
	"github.com/aadog/go-ndk/utils"
)

const (
	RTLD_LAZY     = int(C.RTLD_LAZY)
	RTLD_NOW      = int(C.RTLD_NOW)
	RTLD_GLOBAL   = int(C.RTLD_GLOBAL)
	RTLD_LOCAL    = int(C.RTLD_LOCAL)
	RTLD_NODELETE = int(C.RTLD_NODELETE)
	RTLD_NOLOAD   = int(C.RTLD_NOLOAD)
)

func DlOpen(name string, flag int) unsafe.Pointer {
	return unsafe.Pointer(C.dlopen((*C.char)(utils.CString(name)), C.int(flag)))
}
func DlSym(handle unsafe.Pointer, symbol string) unsafe.Pointer {
	sym := (*C.char)(utils.CString(symbol))
	return unsafe.Pointer(C.dlsym(handle, sym))
}
func DlClose(dl unsafe.Pointer) {
	C.dlclose(dl)
}
func DlError() *string {
	ptr := C.dlerror()
	if ptr == nil {
		return nil
	}
	s := C.GoString(ptr)
	return &s
}
