package p1

import "fmt"

const (
	KB = 1 << (10 * iota)
	MB
	GB
)

func ConstEnum() {

	const (
		A int     = 10
		B float64 = 10.90
		C string  = "hello"
	)

	fmt.Println(A, B, C)
	fmt.Println(KB, MB, GB)
}
