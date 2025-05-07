package channel

import (
	"fmt"
	"time"
)

// NotifyAndMultiplexingCase 通知和多路复用
func NotifyAndMultiplexingCase() {
	ints := make(chan int, 0)
	strings := make(chan string, 0)
	stu := make(chan struct{}, 0)

	go func() {
		for i := 0; i < 10; i++ {
			ints <- i
		}
		close(ints)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			strings <- fmt.Sprintf("字符串 %d", i)
		}
	}()

	f1(ints, strings, stu)
	time.Sleep(3 * time.Second)
	close(stu) // 关闭协程，会向所有监听者发送一个零值数据
}

// select 多三个通道多路复用，作为一个子句整体阻塞，有任意一个通道可用时就不会阻塞，从而执行相应的case代码
func f1(ints chan int, strings chan string, stu chan struct{}) {
	i := 0
	for {
		// case是无序的
		select {
		case a := <-ints:
			fmt.Println(a)
			//case a := <-strings:
			//	fmt.Println(a)
			//case <-stu:
			//	fmt.Println("退出死循环")
			//	return
			//default:
			// 如果加入default，即便三个通道都没有可读取的数据，都在阻塞中，select会执行到default，而不会阻塞
			//fmt.Println("default")
		}
		i++
		fmt.Println("count: ", i)
	}
}
