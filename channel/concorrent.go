package channel

import "fmt"

// ConcurrentCase 并发场景下的同步机制
func ConcurrentCase() {
	ints := make(chan int, 10)

	// 两个并发写入，是没有顺序的
	go func() {
		for i := 0; i < 100; i++ {
			ints <- i
		}
	}()

	go func() {
		for i := 200; i < 300; i++ {
			ints <- i
		}
	}()

	// 单个协程读取，从队列里面获取是有顺序的
	go func() {
		for i := range ints {
			fmt.Println(i)
		}
	}()
}
