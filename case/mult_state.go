package _case

import "fmt"

// 多态 和 泛型 case

// Go 没有显示的实现接口的关键字
// 定义了和接口完全相同的方法，则自动实现了接口

type Person interface {
	eat()
	play()
}

type Teacher struct {
	Name string
	Age  int
}

func (*Teacher) eat() {
	fmt.Println("Teacher eat")
}

func (*Teacher) play() {
	fmt.Println("Teacher play")
}

type Student struct {
}

func (Student) eat() {
	fmt.Println("Student eat")
}

func (Student) play() {
	fmt.Println("Student play")
}

func PersonCase(person Person) {
	person.eat()
	person.play()
}

func GenericsCase(obj interface{}) {
	person, ok := obj.(Person) // 断言是否 Person 类型
	if ok {
		person.eat()
	} else {
		fmt.Println("not Person generic")
	}
}
