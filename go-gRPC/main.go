package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kopher1601/go-studies/go-gRPC/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Koma",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "080-1234-5679"},
		Projects: map[string]*pb.Company_Project{
			"ProjectX": &pb.Company_Project{},
		},
		Profile: &pb.Employee_Text{
			Text: "My name is Koma.",
		},
		Birthday: &pb.Date{
			Day:   1,
			Month: 1,
			Year:  2000,
		},
	}

	// serialize
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	if err := os.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}

	in, err := os.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read", err)
	}

	//var newEmployee pb.Employee
	newEmployee := &pb.Employee{}
	if err := proto.Unmarshal(in, newEmployee); err != nil {
		log.Fatalln("Can't deserialize", err)
	}

	fmt.Println(newEmployee)
}
