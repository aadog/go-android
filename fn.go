package go_java

import "C"
import "unsafe"

func CString(b string) *C.char {
	temp := []byte(b)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return (*C.char)(unsafe.Pointer(&utf8StrArr[0]))
}
