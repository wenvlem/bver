# bver
[![Build Status](https://travis-ci.org/wenvlem/bver.svg?branch=master)](https://travis-ci.org/wenvlem/bver)
<!-- [![GoDoc](https://godoc.org/github.com/wenvlem/bver?status.svg)](https://godoc.org/github.com/wenvlem/bver) -->

bver (pronounced beaver) is a w3c-formatted http access log reader that reports summarized statistics to stdout.

#### Future Improvements
 - [ ] read logs from stdin
 - [ ] output statistics in json or other machine readable format
 - [ ] export statistics via socket to remote server

#### todo
 - [x] add configurable options (log file, report frequency, threshold, duration)
 - [x] create log file if not exist
 - [ ] improve newSaturationMonitor function/options
 - [ ] track down memory leak when spamming logs
 - [ ] implement own file tailing logic
 - [ ] dockerize
