package _case

import "fmt"

// 异常处理 case

func TryCatchCase() {
	// 若不使用 defer ， recover 来恢复程序，则会抛出异常，程序中断
	// defer 按照 先进后出原则，定义多个 defer ，先进的最后执行
	// defer 资源释放 / 异常捕获处理 / 参数预处理
	// defer 执行时机： return 之后，调用方拿到返回值之前
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	MyPanic()
}

func MyPanic() {
	panic("My Panic Q!")
}

func DeferCase() {
	i := 1
	// 传参，i 传入 defer，定义处预存的值
	defer func(a int) {
		fmt.Println("defer a=", a)
	}(i + 1)

	i = 10
	// 闭包，引用局部变量，函数处理完成之后，最后的值
	defer func() {
		i++
		fmt.Println("defer i=", i)
	}()
	i = 99
	return
}

var x int = 1

func DeferCase2() {
	i, i2 := f2()
	fmt.Println("DeferCase2", i, *i2, x)
	//fmt.Println("DeferCase2", i, i2, x) 打印 i2的内存地址
}

func f2() (int, *int) {
	defer func() {
		x = 100 // 修改返回值
	}()
	fmt.Println("f2 x=", x)
	return x, &x
}
