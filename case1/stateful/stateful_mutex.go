package main

import (
	"sync"

	"github.com/fairyhunter13/singleton-multiton/case1/singleton"
)

type wrapper struct {
	mutex         *sync.Mutex
	instanceEmail *singleton.StatefulEmail
}

var (
	wrap *wrapper
	wg   *sync.WaitGroup
)

func init() {
	wrap = &wrapper{
		mutex:         new(sync.Mutex),
		instanceEmail: new(singleton.Email).GetInstanceEmail(),
	}
	wg = new(sync.WaitGroup)
}

func main() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		wrap.sendEmail("test@example.com", "Hi and hello! Test!")
	}()
	wrap.sendEmail("hafiz@example.com", "Hi and hello! Hafiz!")
	wg.Wait()
}

func (w *wrapper) sendEmail(destination, body string) {
	w.mutex.Lock()
	w.instanceEmail.SetTemplate(destination, body)
	w.instanceEmail.SendEmail()
	w.mutex.Unlock()
}
