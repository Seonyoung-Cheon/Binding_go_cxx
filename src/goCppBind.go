package main

//#cgo LDFLAGS: -L/home/cheonsy/lattigo/gpu-core/BindingTest/temp-go-test/build/cpp -lgoCppBind
//#include <stdlib.h>
//#include "/home/cheonsy/lattigo/gpu-core/BindingTest/temp-go-test/cpp/goCppBind.h"
import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
)

//export GoFunction1
func GoFunction1(intValue C.int, charValue *C.char) *C.char {
	fmt.Println("[Go] From C intValue:" + strconv.Itoa(int(intValue)) + " *CharValue:" + C.GoString(charValue))
	returnStr := "Hello"
	return C.CString(returnStr)
}

func main() {
	var cStructure C.STUDENT
	cStructure.name = C.CString("Cheon.Seon.Young")
	cStructure.entYear = C.int(25)
	C.printStudentCpp(&cStructure)
	C.increaseYear(&cStructure)
	C.increaseYear(&cStructure)
	C.free(unsafe.Pointer(cStructure.name))

	var p *C.STUDENT
	name := C.CString("Kim.M.S")
	age := C.int(28)
	p = C.createRecord(name, age)
	C.printStudentCpp(p)
	C.increaseYear(p)
	C.printStudentCpp(p)
	C.free(unsafe.Pointer(name))
	C.free(unsafe.Pointer(p))

	cStr := C.callGoFunction()
	fmt.Println("[Go] call Result=" + C.GoString(cStr))
	// C.free(unsafe.Pointer(cStr))

}
