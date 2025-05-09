package _func

import (
	"log"
)

func Fib(n int) int {
	if n <= 2 {
		log.Fatal("n must > 2")
	}

	t := tool()
	res := 0
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

func tool() func() int {
	x0 := 0
	x1 := 1
	x2 := 0

	// 闭包引用函数的局部变量，导致变量从栈逃逸到堆上
	return func() int {
		x2 = x0 + x1
		x0 = x1
		x1 = x2
		return x2
	}
}
