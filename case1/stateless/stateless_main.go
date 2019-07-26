package main

import (
	"sync"

	"github.com/fairyhunter13/singleton-multiton/case1/singleton"
)

var (
	instance = new(singleton.ReportSingleton).GetInstanceStatelessReport()
	wg       = new(sync.WaitGroup)
)

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		sendReport("test@example.com")
	}()
	go func() {
		defer wg.Done()
		sendReport("hafiz@example.com")
	}()
	sendReport("hello@example.com")
	wg.Wait()
}

func sendReport(dest string) {
	instance.SendTo(dest)
}
