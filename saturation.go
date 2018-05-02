package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type reqCount struct {
	count     int64
	threshold int64
	ttl       time.Duration
}

func init() {
	saturation = newSatMon()
}

var (
	saturation reqCount // saturation will serve as the saturation monitor

	// configurable
	duration = 120 // duration is the size of the monitoring window. Will also serve us as the ttl.
	psLimit  = 10  // psLimit is the threshold for things (requests) per second.

	defTTL = time.Second * time.Duration(duration)
)

// todo: use functional options, set above vars as defaults inside this function
func newSatMon() reqCount {
	return reqCount{count: 0, threshold: int64(psLimit * duration), ttl: defTTL}
}

func (r *reqCount) push(ttl ...time.Duration) {
	var t time.Duration
	if ttl == nil || len(ttl) < 1 {
		t = r.ttl
	} else {
		// use the first element
		t = ttl[0]
	}
	atomic.AddInt64(&r.count, 1)
	go r.pop(t)
}

func (r *reqCount) pop(ttl time.Duration) {
	// todo: add total to be popped and do at once, would need to track their pop "time"
	<-time.After(ttl)
	if r.count > 0 {
		atomic.AddInt64(&r.count, -1)
	}
}

func (r *reqCount) monitor() {
	triggered := false
	for {
		if triggered && atomic.LoadInt64(&r.count) < r.threshold {
			fmt.Printf("High traffic recovered at %s\n", time.Now().Format("15:04:05.1234"))
			triggered = false
		}
		if !triggered && atomic.LoadInt64(&r.count) >= r.threshold {
			fmt.Printf("High traffic generated an alert - hits = %d, triggered at %s\n", r.count, time.Now().Format("15:04:05.1234"))
			triggered = true
		}

		<-time.After(time.Second)
	}
}
