package utils

import (
	"time"
)

type TimingWheel struct {
	ticker  *time.Ticker
	cs      []chan struct{}
	buckets int
	pos     int
}

func NewTimingWheel(interval time.Duration, buckets int) *TimingWheel {
	w := &TimingWheel{
		pos:     0,
		cs:      make([]chan struct{}, buckets),
		buckets: buckets,
		ticker:  time.NewTicker(interval),
	}

	for i := 0; i < buckets; i++ {
		w.cs[i] = make(chan struct{})
	}

	go w.run()
	return w
}

func (w *TimingWheel) After(timeout int) <-chan struct{} {
	return w.cs[(w.pos+timeout)%w.buckets]
}

func (w *TimingWheel) run() {
	for {
		<-w.ticker.C
		w.onTicker()
	}
}

func (w *TimingWheel) onTicker() {
	close(w.cs[w.pos])
	w.cs[w.pos] = make(chan struct{})
	w.pos = (w.pos + 1) % w.buckets
}
