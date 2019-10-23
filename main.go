package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	personpb "github.com/HeWiTo/ex-jsontoprotobuf/protofile/person"
	eyepb "github.com/HeWiTo/ex-jsontoprotobuf/protofile/person"
)

func main() {
	pm := doEnum()

	jsonDemo(pm)
}

func jsonDemo(pb proto.Message) {
	smAsString := toJSON(pb)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct")
	}
}

// func doEnum() {
// 	em := enumpb.EnumMessage{
// 		Id:           42,
// 		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
// 	}

// 	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
// 	fmt.Println(em)
// }

func doEnum() *personpb.PersonMessage {
	pm := personpb.PersonMessage{
		Age:       17,
		FirstName: "Foo",
		LastName:  "Bar",
		EyeColour: personpb.PersonMessage.,
	}
	fmt.Println(pm)

	return &pm
}
