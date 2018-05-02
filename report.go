package main

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type (
	// stats defines collectable stats from a common log formatted entry.
	stats struct {
		requests   reqSlice      // requests is a slice of requests.
		reqTex     *sync.RWMutex // reqTex is requests' lock.
		responses  resSlice      // responses is a slice of responses.
		resTex     *sync.RWMutex // resTex is responses' lock.
		txBytes    int           // txBytes is the total bytes transmitted to the client.
		reportFreq int           // reportFreq is how frequently to print a summary.
	}

	// request defines a countable request.
	request struct {
		count   int    // count is the count of occurrences of the section.
		section string // section is the first part of the path requested. (if path == "/pages/thing", section = "/pages")
	}

	// response defines a countable response.
	response struct {
		count int // count is the count of occurrences of the code.
		code  int // code is the code the server responded with.
	}

	// for implementing the sort interface
	reqSlice []request  // reqSlice is a slice of requests.
	resSlice []response // resSlice is a slice of responses.
)

// buildReport aggregates collected statistics and prints the data when configured.
func buildReport(ctx context.Context, e chan logEntry, s *satMon, reportFreq int) {
	var t = time.Tick(time.Second * time.Duration(reportFreq))

	report := stats{
		reqTex:     &sync.RWMutex{},
		resTex:     &sync.RWMutex{},
		reportFreq: reportFreq,
	}

	go s.monitor(ctx)

	for {
		select {
		case <-t:
			report.print()
			report.clear()
		case entry := <-e:
			s.push()
			report.addRequest(request{section: entry.request.path, count: 1})
			report.addResponse(response{code: entry.respCode, count: 1})
			report.txBytes += entry.txBytes
		case <-ctx.Done():
			return
		}
	}
}

// clear resets the stats' data to 0. (preserves the mutexes)
func (s *stats) clear() {
	s.reqTex.Lock()
	s.resTex.Lock()
	defer s.reqTex.Unlock()
	defer s.resTex.Unlock()
	s.requests = reqSlice{}
	s.responses = resSlice{}
	s.txBytes = 0
}

// print prints the summarized stats.
func (s stats) print() {
	// check whether to print header/footer (each printer has it's own check)
	if s.txBytes == 0 && len(s.responses) == 0 && len(s.requests) == 0 {
		return
	}
	fmt.Println("---------------------------------------")
	s.printRequest()
	s.printResponse()
	s.printTxBytes()
	fmt.Println("=======================================")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE TXBYTE
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// printTxBytes prints the transmitted bytes.
func (s stats) printTxBytes() {
	if s.txBytes != 0 {
		fmt.Printf("Transmitted:\n %dbps\n", s.txBytes/s.reportFreq)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE RESPONSE
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// addResponse increases a response count by 1.
func (s *stats) addResponse(r response) {
	s.resTex.Lock()
	defer s.resTex.Unlock()

	for i := range s.responses {
		if s.responses[i].code == r.code {
			s.responses[i].count++
			return
		}
	}
	s.responses = append(s.responses, r)
}

// printResponse prints the response stats.
func (s stats) printResponse() {
	if len(s.responses) == 0 {
		return
	}
	s.resTex.RLock()
	defer s.resTex.RUnlock()
	sort.Sort(s.responses)
	fmt.Println("Responses:")
	for i := range s.responses {
		fmt.Printf("%3d %d\n", s.responses[i].count, s.responses[i].code)
	}
	fmt.Println()
}

// Sort interface methods

// Len allows resSlice to implement the sort.Interface interface.
func (r resSlice) Len() int {
	return len(r)
}

// Less allows resSlice to implement the sort.Interface interface.
func (r resSlice) Less(i, j int) bool {
	if r[i].count > r[j].count {
		return true
	} else if r[i].count == r[j].count {
		return r[i].code < r[j].code
	}
	return false
}

// Swap allows resSlice to implement the sort.Interface interface.
func (r resSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST REQUEST
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// addRequest increases a request's section count by 1.
func (s *stats) addRequest(r request) {
	s.reqTex.Lock()
	defer s.reqTex.Unlock()

	section := "/"
	if strings.Count(r.section, "/") > 1 {
		section = "/" + strings.FieldsFunc(r.section, func(c rune) bool { return c == '/' })[0]
	}

	r.section = section

	for i := range s.requests {
		if s.requests[i].section == r.section {
			s.requests[i].count++
			return
		}
	}
	s.requests = append(s.requests, r)
}

// printRequest prints the request stats.
func (s stats) printRequest() {
	if len(s.requests) == 0 {
		return
	}
	s.reqTex.RLock()
	defer s.reqTex.RUnlock()
	sort.Sort(s.requests)
	fmt.Println("Requests:")
	for i := range s.requests {
		fmt.Printf("%3d %s\n", s.requests[i].count, s.requests[i].section)
	}
	fmt.Println()
}

// Sort interface methods

// Len allows reqSlice to implement the sort.Interface interface.
func (r reqSlice) Len() int {
	return len(r)
}

// Less allows reqSlice to implement the sort.Interface interface.
func (r reqSlice) Less(i, j int) bool {
	if r[i].count > r[j].count {
		return true
	} else if r[i].count == r[j].count {
		return r[i].section < r[j].section
	}
	return false
}

// Swap allows reqSlice to implement the sort.Interface interface.
func (r reqSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
