package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func scope() {
	var a int
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	var ptr *int = &a
	fmt.Printf("pointer %p,p %d %x,ptr %p\n",pointer,p,p,ptr)
}
func main(){
	// 指针
	scope()
	// 异常
	e := errors.New("divide by zero")
	fmt.Printf("%v\n",e)
	fmt.Printf("%+v\n",e)
	fmt.Printf("%#v\n",e)
}