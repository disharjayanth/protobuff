package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/disharjayanth/protobuff/tree/main/4-golang-protobuff/src/simple/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)

	jsonDemo(sm)
}

func jsonDemo(sm proto.Message) {
	smToString := toJSON(sm)
	fmt.Println("JSON sm:", smToString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smToString, sm2)
	fmt.Println("sm2 PB struct:", sm2)
}

func readAndWriteDemo(sm proto.Message) {
	var sm2 simplepb.SimpleMessage

	writeToFile("simple.bin", sm)
	readFromFile("simple.bin", &sm2)

	fmt.Println("sm2:", sm2)
}

func fromJSON(in string, pb proto.Message) {
	sb := []byte(in)
	if err := protojson.Unmarshal(sb, pb); err != nil {
		log.Fatalln("Error while converting JSON string to PB struct:", err)
	}

}

func toJSON(pb proto.Message) string {
	b, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Error while converting proto to JSON string:", err)
		return err.Error()
	}
	return string(b)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes:", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file:", err)
		return err
	}

	fmt.Println("Date has been written.")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading from file:", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Error putting slice of bytes into protocol buffer message struct:", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
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

	return &sm
}
