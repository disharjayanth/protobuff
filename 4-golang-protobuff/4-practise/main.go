package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/disharjayanth/protobuff/tree/main/4-golang-protobuff/4-practise/proto/addresspb"
	"google.golang.org/protobuf/proto"
)

func main() {
	abm := getAddressBook()
	fmt.Println("AddressBook:", abm)

	fmt.Println("Writing it to disk")
	writeToDisk("serialize.bin", abm)

	abm2 := &addresspb.AddressBook{}
	readFromDisk("serialize.bin", abm2)
	fmt.Println("abm2:", abm2)
}

func writeToDisk(fname string, abm proto.Message) {
	sb, err := proto.Marshal(abm)
	if err != nil {
		log.Fatalln("Cannot serialize addressbook message struct into slice of byte:", err.Error())
	}
	if err := ioutil.WriteFile(fname, sb, 0644); err != nil {
		log.Fatalln("Cannot write serialized data to file:", err.Error())
	}
}

func readFromDisk(fname string, abm2 proto.Message) {
	sb, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Cannot read from file:", err.Error())
	}

	if err := proto.Unmarshal(sb, abm2); err != nil {
		log.Fatalln("Cannot unmarshal slice of byte back to addressbook message struct:", err.Error())
	}
}

func getAddressBook() *addresspb.AddressBook {
	abm := addresspb.AddressBook{
		People: []*addresspb.Person{
			&addresspb.Person{
				Name:  "Ram",
				Id:    20,
				Email: "Ram@gmail.com",
				Phones: []*addresspb.Person_PhoneNumber{
					&addresspb.Person_PhoneNumber{
						Number: "4347343",
						Type:   addresspb.Person_HOME,
					},
					&addresspb.Person_PhoneNumber{
						Number: "8233242834",
						Type:   addresspb.Person_WORK,
					},
				},
			},
			&addresspb.Person{
				Name:  "Shiva",
				Id:    22,
				Email: "Shiva@gmail.com",
				Phones: []*addresspb.Person_PhoneNumber{
					&addresspb.Person_PhoneNumber{
						Number: "23362362",
						Type:   addresspb.Person_MOBILE,
					},
					&addresspb.Person_PhoneNumber{
						Number: "236232",
						Type:   addresspb.Person_HOME,
					},
				},
			},
		},
	}

	return &abm
}
