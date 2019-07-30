package singleton

import (
	"sync"

	"github.com/fairyhunter13/singleton-multiton/case1/singleton"
)

type wrapper struct {
	mutex         *sync.Mutex
	instanceEmail *singleton.StatefulEmail
}

var (
	wrap = &wrapper{
		mutex:         new(sync.Mutex),
		instanceEmail: new(singleton.Email).GetInstanceEmail(),
	}
)

func (w *wrapper) sendEmail(destination, body string) {
	w.mutex.Lock()
	w.instanceEmail.SetTemplate(destination, body)
	w.instanceEmail.SendEmail()
	w.mutex.Unlock()
}
