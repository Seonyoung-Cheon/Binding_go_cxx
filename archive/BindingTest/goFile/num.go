package num

// #include "/home/cheonsy/lattigo/gpu-core/BindingTest/src/include/num.h"
import "C"
import (
	"unsafe"
)

type GoNum struct {
	num C.Num
}

func New() GoNum {
	var ret GoNum
	ret.num = C.NumInit()
	return ret
}
func (n GoNum) Free() {
	C.NumFree((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) Inc() {
	C.NumIncrement((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) GetValue() int {
	return int(C.NumGetValue((C.Num)(unsafe.Pointer(n.num))))
}
