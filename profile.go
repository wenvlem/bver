package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

func init() {
	sigs := make(chan os.Signal)
	go watchSig(sigs)
	signal.Notify(sigs, syscall.SIGHUP)
}

func watchSig(sig chan os.Signal) {
	for {
		select {
		case <-sig:
			f, err := os.Create(fmt.Sprintf("./mem-%s.mprof", time.Now().Format("15:04:05.1234")))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Writing profile...")
			// runtime.GC() //get up to date stats
			pprof.WriteHeapProfile(f)
			f.Close()
			fmt.Println("Profile wrote")
		}
	}
}

/*
// after 434859 logs dumped rapidly to the log file
// `ps` shows ~1.3GB
%MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
 5.4 2111604 1347256 pts/2 Sl+  09:40   0:25 bver -l=/tmp/logs

// pprof `top` shows ~210MB
      flat  flat%   sum%        cum   cum%
  116.04MB 55.08% 55.08%   116.04MB 55.08%  runtime.malg
      64MB 30.38% 85.46%    65.27MB 30.98%  time.NewTimer
   24.50MB 11.63% 97.09%    89.77MB 42.61%  main.(*satMon).pop

// dig into tailer for the missing 1+GB
*/
