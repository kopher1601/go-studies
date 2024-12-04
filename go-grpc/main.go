package main

import (
	"github.com/golang/protobuf/jsonpb"
	"go-grpc/pb"
	//"google.golang.org/protobuf/proto"
	"log"
	//"os"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Kakao",
		Email:       "test@test.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-1234", "090-1234-1234"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Kakao",
		},
		Birthday: &pb.Date{
			Year:  2024,
			Month: 12,
			Day:   4,
		},
	}

	// serialize
	//binData, err := proto.Marshal(employee)
	//if err != nil {
	//	log.Fatalln("Can't serialize", err)
	//}
	//
	//if err := os.WriteFile("test.bin", binData, 0666); err != nil {
	//	log.Fatalln("Can't write", err)
	//}
	//
	//in, err := os.ReadFile("test.bin")
	//if err != nil {
	//	log.Fatalln("Can't read", err)
	//}
	//
	//readEmployee := &pb.Employee{}
	//err = proto.Unmarshal(in, readEmployee)
	//if err != nil {
	//	log.Fatalln("Can't deserialize", err)
	//}
	//
	//fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee) // json stringに変換
	if err != nil {
		log.Fatalln("Can't marshal employee:", err)
	}
	//fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal employee:", err)
	}
}
