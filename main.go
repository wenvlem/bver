// bver is a w3c-formatted http access log reader that reports summarized statistics to stdout.
package main

import (
	"context"
)

func main() {
	logFile := "/tmp/logs"
	outChan := make(chan string)
	entries := make(chan logEntry)

	ctx := context.Background()

	// watch the logfile
	go tail(ctx, logFile, outChan)

	// collect and show statistics
	go buildReport(ctx, entries, newSaturationMonitor(), 2)

	// parse log entries and send to report
	for {
		select {
		case m := <-outChan:
			e, err := parseLine(m)
			if err != nil {
				continue
			}
			entries <- e
		case <-ctx.Done():
			return
		}
	}
}
