package main

import (
	"fmt"

	"github.com/disharjayanth/protobuff/tree/main/4-golang-protobuff/3-complex-proto/proto/complexpb"
)

func main() {
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Message",
			},
			&complexpb.DummyMessage{
				Id:   4,
				Name: "Fourth Message",
			},
		},
	}

	fmt.Println(cm)
}
