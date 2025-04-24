package p1

import (
	_case "first-demo/case"
	"fmt"
)

func NewCase() {
	// new可以创建任意类型，并返回指针，分配内存空间
	// 如果是引用类型，不知道空值是什么，不会进行初始化，因此内存空间的值是 nil
	// 引用类型：切片，map，interface，func，channel，指针类型，其它都是值类型
	mapPtr := new(map[string]_case.User)
	//if mapPtr == nil {
	if *mapPtr == nil {
		fmt.Println("mapPtr is nil")
	}
	//(*mapPtr)["first-demo"] = &_case.User{}

	slicePtr := new([]_case.User)
	if *slicePtr == nil {
		fmt.Println("slicePtr is nil")
	}

	// 切片 虽然是nil，但是可以append
	*slicePtr = append(*slicePtr, _case.User{"k", 100, _case.FEMALE})

	// 如果是值类型，会进行零值初始化，可以直接使用
	strPtr := new(string)
	userPtr := new(_case.User)
	//if *strPtr == nil {
	//	fmt.Println("strPtr is nil")
	//}
	//if *userPtr == nil {
	//	fmt.Println("userPtr is nil")
	//}

	fmt.Println(mapPtr, slicePtr, *strPtr, userPtr)
}

func MakeCase() {
	// make 仅用于切片，map，通道的初始化
	slice := make([]int, 5, 10)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	//append(slice, 100)
	slice = append(slice, 101)
	slice = append(slice, 102)

	slice2 := slice[2:]

	mp := make(map[string]string, 10)
	mp["name"] = "goland"
	mp["love"] = "java"
	mp["interesting"] = "rust"

	// 无序
	for k, v := range mp {
		fmt.Println("k:", k, "v:", v)
	}

	ch := make(chan int, 10)
	ch1 := make(chan<- string, 10)
	ch2 := make(<-chan bool, 5)

	fmt.Println(slice, slice2, mp, ch, ch1, ch2)
}
