package store

import "time"

type Timer interface {
	Sleep(t time.Duration)
	Now() time.Time
}

type RealTimer struct{}

func (*RealTimer) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (*RealTimer) Now() time.Time {
	return time.Now()
}
