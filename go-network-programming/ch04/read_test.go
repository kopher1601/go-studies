package main

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadInfoBuffer(t *testing.T) {
	payload := make([]byte, 1<<24) // 16MB
	_, err := rand.Read(payload)   // 랜덤한 페이로드 생성
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1<<19)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break
		}

		t.Logf("read %d bytes", n) // buf[:n] 은 conn 객체에서 읽은 데이터
	}

	conn.Close()
}
