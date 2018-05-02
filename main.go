// bver is a w3c-formatted http access log reader that reports summarized statistics to stdout.
package main

// import "context"
// import "time"

// import "os"
// import "fmt"
// import "sort"
// import "sync"

// var logs = []string{
// 	`127.0.0.1 - - [01/May/2018 12:29:13] "GET / HTTP/1.1" 200 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:13] code 404, message File not found`,
// 	`127.0.0.1 - - [01/May/2018 12:29:13] "GET /default-styles.css HTTP/1.1" 404 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:21] "GET / HTTP/1.1" 200 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:21] code 404, message File not found`,
// 	`127.0.0.1 - - [01/May/2018 12:29:21] "GET /default-styles.css HTTP/1.1" 404 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] "GET / HTTP/1.1" 200 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] code 404, message File not found`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] "GET /default-styles.css HTTP/1.1" 404 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] "GET / HTTP/1.1" 200 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] code 404, message File not found`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] "GET /default-styles.css HTTP/1.1" 404 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] "GET / HTTP/1.1" 200 -`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] code 404, message File not found`,
// 	`127.0.0.1 - - [01/May/2018 12:29:22] "GET /default-styles.css HTTP/1.1" 404 -`,
// }

func main() {
	// file, err := os.Open("/tmp/logs")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// entries := []logEntry{}
	// 	ctx := context.Background()
	// 	entries := make(chan logEntry,16)

	// 	go buildReport(ctx, entries)

	// keepLogging:
	// 	for i := range logs {
	// 		e, err := parseLine(logs[i])
	// 		if err != nil {
	// 			continue
	// 		}
	// 		entries <- e
	// 	}
	// 	time.Sleep(time.Millisecond * 1500)
	// 	goto keepLogging
}
