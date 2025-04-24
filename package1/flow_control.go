package p1

import "fmt"

func FlowControl() {
	switchCase("b", "hello")
	forCase()
}

func switchCase(name string, in interface{}) {
	switch name {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
		fallthrough
	case "c":
		fmt.Println("c")
	case "d":
		fmt.Println("d")
	}

	switch in.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unknown")
	}

}

func forCase() {
	fmt.Println("myLab1")
myLab1:
	for i := 0; i < 10; i++ {
		fmt.Println("forCase11 ", i)
		if i == 5 {
			break myLab1
		}
	}

	var a = 0
myLab2:
	fmt.Println("myLab2")
	for i := 0; i < 10; i++ {
		fmt.Println("forCase22 ", i)
		a++
		if a == 5 {
			a += 1
			goto myLab2
		}
	}
}
