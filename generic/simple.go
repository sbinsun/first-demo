package generic

import (
	"fmt"
)

// -------------------------------- 泛型函数

func getMaxNum[T int | float64](a, b T) T {
	//func getMaxNum[T interface{int | float64}](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func SampleCase() {
	// 可以自动推断泛型类型，也可以显示指定
	fmt.Println(getMaxNum(32, 10))
	fmt.Println(getMaxNum[float64](32.35, 10.124))
}

// -------------------------------- 定义复杂泛型约束

type CusNumT interface {
	// ~ 表示衍生类型
	// ｜ 表示或者，取并集
	// 多行 表示取交集
	int16 | int32 | ~int64
	int32 | ~int64 | float32 | float64
}

// MyInt64 int64 的衍生类型，是具有int64的新类型，与int64是不同的类型
type MyInt64 int64

// MyInt32 为int32的别名，本质是一样的
type MyInt32 = int32

func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func CusNumCase() {
	var a, b int32 = 1, 2
	fmt.Println(getMaxCusNum(a, b))

	// 不支持的类型
	//var c, d float64 = 1.11, 2.222
	//fmt.Println(getMaxCusNum(c, d))

	var e, f MyInt64 = 31, 25
	fmt.Println(getMaxCusNum(e, f))
}

// -------------------------------- SDK 内置泛型

// any 几乎没有任何约束，和没有一样
func anyType[T any](t T) {
	fmt.Println(t)
}

// comparable 只支持 == ，!=
func compareType[T comparable](a, b T) bool {
	return a == b
}

func BuildInCase() {
	anyType("hello")
	anyType(112312)
	fmt.Println(compareType("hello", "no"))
	fmt.Println(compareType(100, 100))
}
