package ch03

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()

	var d net.Dialer // DialContext는 Dialer의 메서드입니다.
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		// 컨텍스트의 데드라인이 지나기 위해 충분히 긴 시간동안 대기합니다.
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		conn.Close()
		t.Fatal("connection did not time out")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error is not a timeout: %v", err)
		}
	}
	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded; got %v", ctx.Err())
	}
}
