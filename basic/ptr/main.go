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

	var m map[string]int                  // 声明map，指定key和value的数据类型
	m = make(map[string]int)              // 初始化，容量为0
	m = make(map[string]int, 5)           // 初始化，容量为5,建议初始化给一个合适的容量，减少扩容的概率
	m = map[string]int{"语文": 0, "数学": 39} // 初始化直接赋值

	var ch chan int        // 声明
	ch = make(chan int, 8) // 初始化环形队列可容纳8个数据

	ch <- 1 // 往管道里写入（send）数据
	ch <- 2
	ch <- 3
	v := <-ch // 从管道里取走（recv）数据
	v = <-ch

	close(ch) // 遍历前必须关闭管道，禁止写入元素

	// 遍历管道里剩下的元素
	for ele := range ch {
		fmt.Println(ele)
	}

	var u User
	user := &u    // 通过取地址符&得到指针
	user = &User{ // 直接创建结构体指针
		Id: 3, Name: "zcy", addr: "beijing",
	}
	user = new(User) // 通过new()函数实体化一个结构体，并返回其指针
}
