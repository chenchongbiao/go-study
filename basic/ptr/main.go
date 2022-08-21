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
	fmt.Printf("pointer %p,p %d %x,ptr %p\n", pointer, p, p, ptr)
}
func main() {
	// 指针
	scope()
	// 异常
	e := errors.New("divide by zero")
	fmt.Printf("%v\n", e)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%#v\n", e)

	var arr1 [5]int = [5]int{} // 数组要指定长度和类型，且长度和类型指定后不可改变
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}            // 给前2个元素复制
	var arr4 = [5]int{2: 15, 4: 30}    // 指定index赋值
	var arr5 = [...]int{3, 2, 6, 5, 4} // 根据{}的元素个数推断出数组的长度
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}} // 数组的元素类型由匿名结构体给定

	// 5行3列，只给前2行赋值，且前2行的所有列没有赋满
	var arr1 = [5][3]int{{1}, {2, 3}}
	// 第1维可以用...推测，第二维不能用...
	var arr2 = [...][3]int{{1}, {2, 3}}

	// 切片
	var s []int              //切片声明，len=cap=0
	s = []int{}              //初始化，len=cap=0
	s = make([]int, 3)       // 初始化，len=cap=3
	s = make([]int, 3, 5)    // 初始化，len=3,cap=5
	s = []int{1, 2, 3, 4, 5} // 初始化len=cap=5
	s2d := [][]int{
		{1}, {2, 3}, //二维数组各行列数相等，但是二维切片各行的len可以不相等
	}
}
