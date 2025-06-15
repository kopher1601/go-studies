package ch03

import (
	"context"
	"io"
	"time"
)

const defaultPingInterval = 30 * time.Second

// go routine 에서 동작하도록 설계됨
func Pinger(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
	case interval = <-reset: // reset 채널에서 초기 간격을 받아 옴
	default:
	}

	if interval <= 0 {
		interval = defaultPingInterval
	}

	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newInterval := <-reset:
			if !timer.Stop() {
				<-timer.C
			}
			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C:
			if _, err := w.Write([]byte("ping")); err != nil {
				// 여기서 연속을호 발생하는 타임아웃을 추적하고 처리함
				return
			}
		}
	}

	_ = timer.Reset(interval)
}
