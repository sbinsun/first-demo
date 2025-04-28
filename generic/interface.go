package generic

import "fmt"

// MyToString 普通接口，可用于变量定义
type MyToString interface {
	toString() string
}

// var a MyToString

func (u MyUser) toString() string {
	return fmt.Sprintf("ID: %d, name: %s, Age: %d", u.Id, u.Name, u.Age)
}

func (u MyAddress) toString() string {
	return fmt.Sprintf("ID: %d, Street: %s, City: %s", u.Id, u.Street, u.City)
}

// GetKey 泛型接口，不可用于变量定义
type GetKey[K comparable] interface {
	Get() K
}

func (u MyUser) Get() int64 {
	return u.Id
}

func (u MyAddress) Get() int32 {
	return u.Id
}

func listToMap[K comparable, V GetKey[K]](list []V) map[K]V {
	mp := make(map[K]V, len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

func InterfaceCase() {
	users := []GetKey[int64]{
		MyUser{
			Id:   1,
			Name: "John Doe",
			Age:  20,
		},
		MyUser{
			Id:   2,
			Name: "go Doe",
			Age:  29,
		}}

	address := []GetKey[int32]{
		MyAddress{
			Id:     1,
			Street: "xi an",
			City:   "xi an",
		},
		MyAddress{
			Id:     2,
			Street: "gua zhou",
			City:   "gua zhou",
		}}

	uMap := listToMap(users)
	aMap := listToMap(address)

	fmt.Println("listToMap:")
	fmt.Println(uMap)
	fmt.Println(aMap)
}
