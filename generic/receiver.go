package generic

import "fmt"

type TheStruct[T interface{ *int32 | *string }] struct {
	Name string
	Data T
}

func (u TheStruct[T]) GetData() T {
	return u.Data
}

func ReceiverCase() {
	var age int32 = 2333
	theStruct := TheStruct[*int32]{
		Name: "monkey",
		Data: &age,
	}

	fmt.Println(*theStruct.GetData())
}
