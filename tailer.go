package main

import (
	"context"
	"time"

	// Don't re-invent the wheel. Plus, shows ability to integrate with libs.
	"github.com/DataDog/datadog-agent/pkg/logs/auditor"
	"github.com/DataDog/datadog-agent/pkg/logs/config"
	"github.com/DataDog/datadog-agent/pkg/logs/input/tailer"
	"github.com/DataDog/datadog-agent/pkg/logs/pipeline/mock"
)

var (
	outChan = make(chan string)
)

// todo: reimplement with fsnotify and own file tailing logic
func tail(ctx context.Context, logFile string) {
	source := config.NewLogSource("access", &config.LogsConfig{Type: "file", Path: logFile})
	sleepDuration := time.Millisecond * 100

	pp := mock.NewMockProvider()
	filesScanner := tailer.New([]*config.LogSource{source}, 4, pp, auditor.New(nil, ""), sleepDuration)
	filesScanner.Start()
	for {
		select {
		case m := <-pp.NextPipelineChan():
			outChan <- string(m.Content())
		case <-ctx.Done():
			filesScanner.Stop()
			return
		}
	}
}
