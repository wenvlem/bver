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
	// thing:= reqCount{count: 0, threshold: int64(2)}
	// go thing.monitor()

	// go thing.push(time.Millisecond*500)
	// go thing.push(time.Second)
	// <-time.After(time.Millisecond*6)
	// fmt.Println(thing)
	// <-time.After(time.Millisecond*600)
	// fmt.Println(thing)
	// <-time.After(time.Millisecond*600)
	// fmt.Println(thing)
	// <-time.After(time.Millisecond*1000)

	// file, err := os.Open("/tmp/logs")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// entries := []logEntry{}

	// 	saturation.threshold = int64(200)
	// 	go saturation.monitor()

	// 	ctx := context.Background()
	// 	entries := make(chan logEntry,16)
	// 	reportFreq = 30
	// 	go buildReport(ctx, entries)

	// 	delay := 500
	// keepLogging:
	// 	for i := range logs {
	// 		e, err := parseLine(logs[i])
	// 		if err != nil {
	// 			continue
	// 		}
	// 		entries <- e
	// 	}
	// 	delay+= 200
	// 	time.Sleep(time.Millisecond * time.Duration(delay))
	// 	goto keepLogging
}
