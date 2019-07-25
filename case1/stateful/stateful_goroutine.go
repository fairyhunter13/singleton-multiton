package main

import (
	"sync"

	"github.com/fairyhunter13/singleton-multiton/case1/singleton"
)

var (
	singletonEmail = new(singleton.Email)
	instanceEmail  = singletonEmail.GetInstanceEmail()
	wg             = new(sync.WaitGroup)
)

func main() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		sendEmail("test@example.com", "Hi and hello! Test!")
	}()
	sendEmail("hafiz@example.com", "Hi and hello! Hafiz!")
	wg.Wait()
}

func sendEmail(destination, body string) {
	instanceEmail.SetTemplate(destination, body)
	instanceEmail.SendEmail()
}
