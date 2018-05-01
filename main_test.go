package main

import (
	"testing"
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

func TestMainFunc(t *testing.T) {
	main()
}
