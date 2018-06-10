package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	pb "go_code/code_db/proto_buf/tutorial/my_protobuf"
	"io/ioutil"
	"log"
	//"reflect"
)

func main() {
	fname := "./data"
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{}
	book.People = append(book.People, &p)
	fmt.Printf("original: %#v\n", book)
	fmt.Printf("original: %#v\n", book.People)

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	restored_book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, restored_book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("restored: %#v\n", restored_book)
	fmt.Printf("restored: %#v\n", restored_book.People)

	fmt.Println(restored_book.People[0].Name)
	fmt.Println(restored_book.People[0].Id)
	fmt.Println(restored_book.People[0].Email)
	fmt.Println(restored_book.People[0].Phones)

	//fmt.Printf("restroed == original? %t\n", restored_book == book)

	// lsl-test: message interface
	p2 := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	fmt.Printf("p == p2 ? %t\n", proto.Equal(&p, &p2))

}
