package generic

import "fmt"

// -------------------------------- 原始玩法

type MyUser struct {
	Id   int64
	Name string
	Age  int
}

type MyAddress struct {
	Id     int32
	Street string
	City   string
}

func mapToList[k comparable, T any](mp map[k]T) []T {
	ts := make([]T, len(mp))
	i := 0
	for _, v := range mp {
		ts[i] = v
		i++
	}
	return ts
}

func myPrint[T any](a chan T) {
	for data := range a {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(map[int64]MyUser, 10)
	userMp[0] = MyUser{
		Id:   1,
		Name: "John Doe",
		Age:  20,
	}
	userMp[1] = MyUser{
		Id:   2,
		Name: "go Doe",
		Age:  29,
	}

	userList := mapToList(userMp)
	uch := make(chan MyUser)
	go myPrint(uch)
	for _, u := range userList {
		uch <- u
	}

	assMp := make(map[int32]MyAddress, 10)
	assMp[1] = MyAddress{
		Id:     1,
		Street: "xi an",
		City:   "xi an",
	}
	assMp[2] = MyAddress{
		Id:     2,
		Street: "gua zhou",
		City:   "gua zhou",
	}

	assList := mapToList(assMp)
	ach := make(chan MyAddress)
	go myPrint(ach)
	for _, a := range assList {
		ach <- a
	}
}

// -------------------------------- 泛型类型

// List 泛型切片
type List[T any] []T

// Map 泛型集合
type Map[K comparable, T any] map[K]T

// Chan 泛型通道
type Chan[T any] chan T

func TTypeCase1() {
	userMp := make(Map[int64, MyUser], 10)
	userMp[0] = MyUser{
		Id:   1,
		Name: "John Doe",
		Age:  20,
	}
	userMp[1] = MyUser{
		Id:   2,
		Name: "go Doe",
		Age:  29,
	}

	var userList List[MyUser]
	userList = mapToList(userMp)
	uch := make(Chan[MyUser])
	go myPrint(uch)
	for _, u := range userList {
		uch <- u
	}

	assMp := make(Map[int32, MyAddress], 10)
	assMp[1] = MyAddress{
		Id:     1,
		Street: "xi an",
		City:   "xi an",
	}
	assMp[2] = MyAddress{
		Id:     2,
		Street: "gua zhou",
		City:   "gua zhou",
	}

	var assList List[MyAddress]
	assList = mapToList(assMp)
	ach := make(Chan[MyAddress])
	go myPrint(ach)
	for _, a := range assList {
		ach <- a
	}
}
