package main

import (
	"context"
	"sort"
	"testing"
	"time"
)

type testCase struct {
	input  string
	output string
	err    bool
}

var (
	logLine = `127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`
	okLine  = `127.0.0.1 - - [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`
	badLine = `something broke`
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestParse(t *testing.T) {
	cases := []testCase{
		testCase{
			input: logLine,
		},
		testCase{
			input: okLine,
		},
		testCase{
			input: badLine,
			err:   true,
		},
	}

	for i := range cases {
		_, err := parseLine(cases[i].input)
		if (err != nil) != cases[i].err {
			t.Errorf("Failed to fail")
			continue
		}
	}
}

func TestAtoi(t *testing.T) {
	i := atoi("hola")
	if i != 0 {
		t.Errorf("Failed to return default value on fail")
		t.FailNow()
	}
}

// Reports

var logs = []string{
	`127.0.0.1 - - [01/May/2018 12:29:13] "GET / HTTP/1.1" 200 -`,
	`127.0.0.1 - - [01/May/2018 12:29:13] code 404, message File not found`,
	`127.0.0.1 - - [01/May/2018 12:29:13] "GET /default-styles.css HTTP/1.1" 404 -`,
	`127.0.0.1 - - [01/May/2018 12:29:21] "GET / HTTP/1.1" 200 -`,
	`127.0.0.1 - - [01/May/2018 12:29:21] code 404, message File not found`,
	`127.0.0.1 - - [01/May/2018 12:29:21] "GET /default-styles.css HTTP/1.1" 404 -`,
	`127.0.0.1 - - [01/May/2018 12:29:22] "GET / HTTP/1.1" 200 -`,
	`127.0.0.1 - - [01/May/2018 12:29:22] code 404, message File not found`,
	`127.0.0.1 - - [01/May/2018 12:29:22] "GET /default-styles.css HTTP/1.1" 404 -`,
	`127.0.0.1 - - [01/May/2018 12:29:22] "GET / HTTP/1.1" 200 -`,
	`127.0.0.1 - - [01/May/2018 12:29:22] code 404, message File not found`,
	`127.0.0.1 - - [01/May/2018 12:29:22] "GET /default-styles.css HTTP/1.1" 404 -`,
	`127.0.0.1 - - [01/May/2018 12:29:22] "GET / HTTP/1.1" 200 -`,
	`127.0.0.1 - - [01/May/2018 12:29:22] code 404, message File not found`,
	`127.0.0.1 - - [01/May/2018 12:29:22] "GET /default-styles.css HTTP/1.1" 404 -`,
	`83.149.9.216 - - [17/May/2015:10:05:03 +0000] "GET /presentations/logstash-monitorama-2013/images/kibana-search.png HTTP/1.1" 200 203023 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:43 +0000] "GET /presentations/logstash-monitorama-2013/images/kibana-dashboard3.png HTTP/1.1" 200 171717 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:47 +0000] "GET /presentations/logstash-monitorama-2013/plugin/highlight/highlight.js HTTP/1.1" 200 26185 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:12 +0000] "GET /presentations/logstash-monitorama-2013/plugin/zoom-js/zoom.js HTTP/1.1" 200 7697 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:07 +0000] "GET /presentations/logstash-monitorama-2013/plugin/notes/notes.js HTTP/1.1" 200 2892 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:34 +0000] "GET /presentations/logstash-monitorama-2013/images/sad-medic.png HTTP/1.1" 200 430406 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:57 +0000] "GET /presentations/logstash-monitorama-2013/css/fonts/Roboto-Bold.ttf HTTP/1.1" 200 38720 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:50 +0000] "GET /presentations/logstash-monitorama-2013/css/fonts/Roboto-Regular.ttf HTTP/1.1" 200 41820 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:24 +0000] "GET /presentations/logstash-monitorama-2013/images/frontend-response-codes.png HTTP/1.1" 200 52878 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:50 +0000] "GET /presentations/logstash-monitorama-2013/images/kibana-dashboard.png HTTP/1.1" 200 321631 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:46 +0000] "GET /presentations/logstash-monitorama-2013/images/Dreamhost_logo.svg HTTP/1.1" 200 2126 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:11 +0000] "GET /presentations/logstash-monitorama-2013/images/kibana-dashboard2.png HTTP/1.1" 200 394967 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
	`83.149.9.216 - - [17/May/2015:10:05:19 +0000] "GET /presentations/logstash-monitorama-2013/images/apache-icon.gif HTTP/1.1" 200 8095 "http://semicomplete.com/presentations/logstash-monitorama-2013/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36"`,
}

func TestReport(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	entries := make(chan logEntry)

	go buildReport(ctx, entries)

	for i := range logs {
		e, err := parseLine(logs[i])
		if err != nil {
			continue
		}
		entries <- e
	}
	time.Sleep(time.Second * 3)
}

func TestSortPrint(t *testing.T) {
	s := stats{}
	s.print()
	s.printRequest()
	s.printResponse()
	s.printTxBytes()

	r := resSlice{
		response{count: 2, code: 202},
		response{count: 2, code: 200},
		response{count: 3, code: 400},
		response{count: 5, code: 301},
		response{count: 1, code: 204},
		response{count: 4, code: 500},
	}
	sort.Sort(r)
	q := reqSlice{
		request{count: 2, section: "/eat"},
		request{count: 2, section: "/popcorn"},
		request{count: 3, section: "/home"},
		request{count: 5, section: "/blog"},
		request{count: 1, section: "/admin"},
		request{count: 4, section: "/wp_admin"},
	}
	sort.Sort(q)
}
