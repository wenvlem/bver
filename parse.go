package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type (
	requestEntry struct {
		method   string
		path     string
		httpType string
	}

	logEntry struct {
		remoteHost string
		userId     string
		authUser   string
		date       string
		request    requestEntry
		respCode   int
		txBytes    int
	}
)

// logRegex is the regex to match the common logfile format defined at
// https://www.w3.org/Daemon/User/Config/Logging.html#common-logfile-format.
// spec:    remotehost rfc931 authuser [date] "request" status bytes
// example: 127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
//
// todo: define and try multiple parsers for different date formats
// 01/May/2018 12:29:22
// 10/Oct/2000:13:55:36 -0700 - (\d{2}/\w{3}/\d{2}(?:\d{2}:){3}\d{2} [-+]\d{4})
var logRegex = regexp.MustCompile( // ^(\S+)\s(\S+)\s(\S+)\s\[(.*)\]\s\"(\S+)\s(\S+)\s(\S+)\"\s(\d{3})\s(\d+)
	`^(\S+)\s` + // remoteHost
		`(\S+)\s` + // userId
		`(\S+)\s` + // authUser
		`\[(.*)\]\s` + // date
		`\"(\S+)\s` + // request.method
		`(\S+)\s` + // request.path
		`(\S+)\"\s` + // request.httpType
		`(\S+)\s` + // respCode
		`(\S+)` + // txBytes
		`.*`)

// parseLine parses a log line and returns a logEntry for further processing.
func parseLine(s string) (logEntry, error) {
	parts := logRegex.FindStringSubmatch(s)
	if len(parts) < 1 {
		return logEntry{}, fmt.Errorf("No match found")
	}

	entry := logEntry{
		remoteHost: parts[1],
		userId:     parts[2],
		authUser:   parts[3],
		date:       parts[4],
		request: requestEntry{
			method:   parts[5],
			path:     parts[6],
			httpType: parts[7],
		},
		respCode: atoi(parts[8]),
		txBytes:  atoi(parts[9]),
	}

	return entry, nil
}

// atoi parses a string and returns an int (0 if there was an error).
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
