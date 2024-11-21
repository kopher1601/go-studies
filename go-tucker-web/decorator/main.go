package main

import (
	"fmt"
	"go-tucker-web/decorator/cipher"
	"go-tucker-web/decorator/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

type SendComponent struct {
}

func (s *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (z *ZipComponent) Operator(data string) {
	// 압축 작업
	ziped, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	// 필드에 가지고 있는 Operator 작업
	z.com.Operator(string(ziped))
}

type EncryptComponent struct {
	key string
	com Component
}

func (e *EncryptComponent) Operator(data string) {
	// 암호 작업
	encrypt, err := cipher.Encrypt([]byte(data), e.key)
	if err != nil {
		panic(err)
	}

	// 필드에 가지고 있는 Operator 작업
	e.com.Operator(string(encrypt))
}

type DecryptComponent struct {
	key string
	com Component
}

func (d *DecryptComponent) Operator(data string) {
	decrypt, err := cipher.Decrypt([]byte(data), d.key)
	if err != nil {
		panic(err)
	}
	d.com.Operator(string(decrypt))
}

type UnzipComponent struct {
	com Component
}

func (u *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	u.com.Operator(string(unzipData))
}

type ReadComponent struct {
}

func (r *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World!")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}
	receiver.Operator(sentData)
	fmt.Println(recvData)

}
