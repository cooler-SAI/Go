package main

import (
	"buffers/example"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {

	person := &example.Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@example.com",
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}

	fmt.Println("Serialized data: ", data)

	newPerson := &example.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}

	fmt.Printf("Deserialized Person: %+v\n", newPerson)
}
