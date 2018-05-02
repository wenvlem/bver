package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// satMon defines a saturation monitor.
type satMon struct {
	count     int64         // count is the number of occurrences.
	threshold int64         // threshold is the upper limit before an alert is "triggerred."
	ttl       time.Duration // ttl is the size of the window the threshold applies to.
}

// newSaturationMonitor returns a pointer to a new satMon.
// todo: use functional options, set above vars as defaults inside this function
func newSaturationMonitor() *satMon {
	s := &satMon{
		count:     0,
		threshold: int64(psLimit * duration),
		ttl:       time.Second * time.Duration(duration),
	}
	return s
}

// push adds 1 to a satMon's counter and calls for it to be subtracted at ttl.
func (r *satMon) push(ttl ...time.Duration) {
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

// pop subtracts 1 to a satMon's counter at ttl.
func (r *satMon) pop(ttl time.Duration) {
	// todo: add total to be popped and do at once, would need to track their pop "time"
	<-time.After(ttl)
	if r.count > 0 {
		atomic.AddInt64(&r.count, -1)
	}
}

// monitor watches a satMon's count, alerting if it exceeds the threshold.
func (r *satMon) monitor(ctx context.Context) {
	triggered := false
	for {
		select {
		default:
			if triggered && atomic.LoadInt64(&r.count) < r.threshold {
				fmt.Printf("High traffic recovered at %s\n", time.Now().Format("15:04:05.1234"))
				triggered = false
			}
			if !triggered && atomic.LoadInt64(&r.count) >= r.threshold {
				fmt.Printf("High traffic generated an alert - hits = %d, triggered at %s\n", r.count, time.Now().Format("15:04:05.1234"))
				triggered = true
			}

			<-time.After(time.Second)
		case <-ctx.Done():
			return
		}
	}
}
