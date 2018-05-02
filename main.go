// bver is a w3c-formatted http access log reader that reports summarized statistics to stdout.
package main

import (
	"context"
	"flag"
	"os"
)

var (
	// configurable options
	logSource       string // logSource is the location of the log to watch and analyze.
	reportFrequency int    // reportFrequency is how frequent a summary will be printed to the screen.
	psLimit         int    // psLimit is the threshold for things (requests) per second.
	duration        int    // duration is the size of the monitoring window. Will also serve us as the ttl.
)

func init() {
	flag.StringVar(&logSource, "l", "/var/log/access.log", "Log location to watch and analyze.")
	flag.IntVar(&reportFrequency, "f", 10, "Frequency at which to print summary.")
	flag.IntVar(&psLimit, "t", 10, "Number of requests per second before printing an alert.")
	flag.IntVar(&duration, "d", 120, "Duration of window to average requests per second.")
	flag.Parse()

	sanitizeOpts()
}

// sanitizeOpts resets sane defaults if bad input is given. It also attempts to create the log
// file if it doesn't already exist.
func sanitizeOpts() {
	if psLimit < 1 {
		psLimit = 10
	}
	if duration < 1 {
		duration = 120
	}
	if reportFrequency < 1 {
		reportFrequency = 10
	}
	if _, err := os.Stat(logSource); os.IsNotExist(err) {
		// ignore error since tailer retries
		os.Create(logSource)
	}
}

func main() {
	outChan := make(chan string)
	entries := make(chan logEntry)

	ctx := context.Background()

	// watch the logfile
	go tail(ctx, logSource, outChan)

	// collect and show statistics
	go buildReport(ctx, entries, newSaturationMonitor(), reportFrequency)

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
