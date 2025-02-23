package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	RequestAndResponse(conn, "request0")
	RequestAndResponse(conn, "request1")
	RequestAndResponse(conn, "request2")

	RequestAndResponse(conn, "close")
	time.Sleep(time.Hour)
}

func RequestAndResponse(conn net.Conn, requestData string) {
	requestByteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write(requestByteData)

	responseData := make([]byte, 100)
	n, err := conn.Read(responseData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(responseData[:n]))
}
