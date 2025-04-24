package p1

import "errors"

// VarSum
// a, b 是形参
// a 值传递
// b 引用传递
// 返回值可以定义多个，以及名称
// 返回值若是值类型，可以默认返回零值
// 返回值若是引用类型，没有初始化，不能直接使用
func VarSum(a int, b *int) (sum int, myError error) {
	if a < 0 && *b < 0 {
		myError = errors.New("a and b is less than 0")
		return
	}

	sum = a + *b
	//return sum, myError
	return
}

// 全局变量
// 小写，当前 package 共有的变量，外部不能当问
// 大写，当前 package 共有的变量，外部能当问
var g int
var G int

type Apple struct {
	Color string
	Size  int
}

func NewApple(color string, size int) *Apple {
	return &Apple{
		Color: color,
		Size:  size,
	}
}

func (apple *Apple) GetColor() string {
	return apple.Color
}

func (apple *Apple) GetSize() int {
	return apple.Size
}
