package p1

import "fmt"

func VarDeclare() {

	var a, b int
	var c int = 100
	var d float64 = 98.21

	// 只适用局部变量
	e := false

	var arr1 = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3}
	var arr3 = [5]int{}

	fmt.Println(a, b, c, d, e, arr1, arr2, arr3)

	var intPtr *int
	var floatPtr *float64

	var f = 100
	add(&f)
	fmt.Println(intPtr, floatPtr, f)

	// interface 可以接收任何类型
	var inter interface{}
	inter = 100
	inter = "inter"
	fmt.Println(inter)
}

func add(x *int) {
	*x = *x + 1
}
