package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	elliot := &Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())

}
