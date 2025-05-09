package _func

import (
	"errors"
	"fmt"
	"log"
)

func sum(a, b int) (sum int, error error) {
	if a < 0 && b < 0 {
		error = errors.New("a and b is less than 0")
		return
	}
	sum = a + b
	return
}

func FuncCase() {
	fmt.Println(sum(-1, -2))

	f1 := sum
	fmt.Println(f1(1, 2))

	f2 := logMiddleware(sum)
	fmt.Println(f2(12, 32))

	f3 := logMiddleware(f2)
	fmt.Println(f3(48, 68))
}

// MySum 将函数作为类型
type MySum func(a, b int) (int, error)

// 将函数作为输入输出，实现中间件，类似方法重写
func logMiddleware(in MySum) MySum {
	return func(a, b int) (int, error) {
		log.Printf("start log a=%d, b=%d", a, b)
		return in(a, b)
	}
}
