package main

import (
	"fmt"

	"github.com/disharjayanth/protobuff/tree/main/4-golang-protobuff/src/simple/simplepb"
)

func main() {
	fmt.Println("Hello world!")
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you!"

	fmt.Println(sm)

	fmt.Println("Id:", sm.GetId())
}
