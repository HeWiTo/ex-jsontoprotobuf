package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/HeWiTo/ex-jsontoprotobuf/protofile"
)

func main() {
	pm := doEnum()

	jsonDemo(pm)
}

func jsonDemo(pb proto.Message) {
	pmAsString := toJSON(pb)
	fmt.Println(pmAsString)

	pm2 := &protofile.PersonMessage{}
	fromJSON(pmAsString, pm2)
	fmt.Println("Successfully created proto struct: ", pm2)
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

func doEnum() *protofile.PersonMessage {
	pm := protofile.PersonMessage{
		Age:       17,
		FirstName: "Foo",
		LastName:  "Bar",
		EyeColour: &protofile.Eye{
			EyeColour: 2,
		},
	}
	fmt.Println(pm)

	return &pm
}
