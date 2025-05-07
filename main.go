package main

import (
	"context"
	_case "first-demo/case"
	channel "first-demo/channel"
	p2 "first-demo/generic"
	p1 "first-demo/package1"
	"fmt"
	"os"
	"os/signal"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	_case.UserCase()

	//say() 调用了同一个包的函数，需要 go build main.go  fun.go ，go run main.go  fun.go

	user := _case.User{}

	t := &_case.Teacher{} // 指针包含 receiver = 值类型/指针类型的所有方法
	//t := _case.Teacher{} // 值只包含 receiver = 值类型方法
	s := _case.Student{}
	_case.PersonCase(t)
	_case.PersonCase(s)

	fmt.Println("===")
	_case.GenericsCase(t)
	_case.GenericsCase(user)

	fmt.Println("===")
	_case.TryCatchCase()
	_case.DeferCase()

	fmt.Println("===")
	_case.DeferCase2()

	// 交叉编译，需要修改三个编译参数
	// GOOS 目标操作系统 darwin freebsd linux windows
	// GOARCH amd64 arm
	// CGO_ENABLE 交叉编译不支持，要设置为 0
	// GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build main.go
	// GOOS=windows GOARCH=amd64 CGO_ENABLE=0 go build main.go

	fmt.Println("===")
	p1.F2()

	fmt.Println("===")
	p1.VarDeclare()

	fmt.Println("===")
	p1.ConstEnum()

	fmt.Println("===")
	p1.NewCase()

	fmt.Println("MakeCase===")
	p1.MakeCase()

	fmt.Println("VarSum===")

	// 实参
	a := 10
	b := 20
	fmt.Println(p1.VarSum(a, &b))
	fmt.Println(a, b)

	fmt.Println(p1.G)
	apple := p1.NewApple("yellow", 12)
	fmt.Println(apple.GetSize())
	fmt.Println(apple.GetColor())

	fmt.Println("FlowControl===")
	p1.FlowControl()

	fmt.Println("Generic===")
	p2.SampleCase()
	p2.CusNumCase()
	p2.BuildInCase()
	p2.TTypeCase()
	p2.TTypeCase1()
	p2.InterfaceCase()
	p2.ReceiverCase()

	fmt.Println("channel===")
	//channel.CommunicationCase()
	channel.NotifyAndMultiplexingCase()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
