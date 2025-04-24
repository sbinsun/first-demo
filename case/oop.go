package _case

import "fmt"

// 面向对象 case
const (
	MALE = 1 << 10 * iota
	FEMALE
	_
	UNKNOWN
)

type Gender uint16

type User struct {
	Name   string
	Age    int
	Gender Gender
}

func (*User) run() { // receiver 为指针类型的方法，值类型的实例本不可以直接调用，编译器做了自动转换
	fmt.Println("user run")
}

func (User) sleep() { // receiver 为值类型的方法
	fmt.Println("user sleep")
}

func UserCase() {
	//u1 := new(User)	 	// 返回指针类型
	//u2 := &User{}		// 返回指针类型
	u := User{ // 返回值类型
		Name:   "gopher",
		Age:    20,
		Gender: FEMALE,
	}
	u.run()
	u.sleep()
	fmt.Println(u.Gender)

	fmt.Println("")
}
