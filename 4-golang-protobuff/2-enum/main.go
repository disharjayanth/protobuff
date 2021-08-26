package main

import (
	"fmt"

	"github.com/disharjayanth/protobuff/tree/main/4-golang-protobuff/2-enum/src/src/enumpb"
)

func main() {
	doEnum()
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_SATURDAY
	fmt.Println(em)
}
