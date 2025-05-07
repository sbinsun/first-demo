package channel

import "fmt"

// CommunicationCase 通过协程共享内存
func CommunicationCase() {
	// 可读可写的通道，设置只有一个容量
	// 通道满时，不能继续写，写协程会阻塞，等待元素读取
	// 没有元素时，读协程不能继续读取，会阻塞，等待元素写入
	ints := make(chan int, 0)
	go writeChan(ints)
	go readChan(ints)
}

// 只写通道
func writeChan(ch chan<- int) {
	for i := 0; i < 30; i++ {
		ch <- i
	}
}

// 制度通道
func readChan(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
