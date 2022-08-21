package main

import (
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
	scope()
}