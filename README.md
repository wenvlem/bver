# bver
[![Build Status](https://travis-ci.org/wenvlem/bver.svg?branch=master)](https://travis-ci.org/wenvlem/bver)
<!-- [![GoDoc](https://godoc.org/github.com/wenvlem/bver?status.svg)](https://godoc.org/github.com/wenvlem/bver) -->

bver (pronounced beaver) is a w3c-formatted http access log reader that reports summarized statistics to stdout.  

#### Usage
```
Usage of bver:
  -d int
    	Duration of window in which to average requests per second. (default 120)
  -f int
    	Frequency at which to print summary (seconds). (default 10)
  -l string
    	Log location to watch and analyze. (default "/var/log/access.log")
  -t int
    	Number of requests per second before printing an alert. (default 10)
```

Example Use:  
```
$ bver -l=/tmp/logs
1525474956711292289 [Info] Opening /tmp/logs
---------------------------------------
Requests:
 15 /

Responses:
 15 200

=======================================
```

#### Future Improvements
 - [ ] read logs from stdin
 - [ ] output statistics in json or other machine readable format
 - [ ] export statistics via socket to remote server
 - [ ] implement own file tailing logic

#### todo
 - [x] add configurable options (log file, report frequency, threshold, duration)
 - [x] create log file if not exist
 - [x] improve newSaturationMonitor function/options
 - [x] dockerize
 - [x] track down memory leak when spamming logs (seems to be somewhere in tailer/ceelog)

# Problem:

## HTTP log monitoring console program

 - [x] Create a simple console program that monitors HTTP traffic on your machine:
 - [x] Consume an actively written-to w3c-formatted HTTP access log (https://en.wikipedia.org/wiki/Common_Log_Format). It should default to reading /var/log/access.log and be overrideable.
 - [x] Display stats every 10s about the traffic during those 10s: the sections of the web site with the most hits, as well as interesting summary statistics on the traffic as a whole. A section is defined as being what's before the second '/' in a URL. For example, the section for "http://my.site.com/pages/create" is "http://my.site.com/pages".
 - [x] Make sure a user can keep the console app running and monitor traffic on their machine.
 - [x] Whenever total traffic for the past 2 minutes exceeds a certain number on average, add a message saying that “High traffic generated an alert - hits = {value}, triggered at {time}”. The default threshold should be 10 requests per second, and should be overrideable.
 - [x] Whenever the total traffic drops again below that value on average for the past 2 minutes, add another message detailing when the alert recovered.
 - [x] Make sure all messages showing when alerting thresholds are crossed remain visible on the page for historical reasons.
 - [x] Write a test for the alerting logic.
 - [x] Explain how you’d improve on this application design.
 - [x] If you have access to a linux docker environment, we'd love to be able to docker build and run your project! If you don't though, don't sweat it. As an example:
```
FROM python:3
RUN touch /var/log/access.log  # since the program will read this by default
WORKDIR /usr/src
ADD . /usr/src
ENTRYPOINT ["python", "main.py"]
```
and we'll have something else write to that log file.
