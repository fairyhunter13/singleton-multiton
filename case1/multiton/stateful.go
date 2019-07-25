package multiton

import (
	"fmt"
	"sync"
	"time"
)

// This file is in another package.

// StatefulEmail is a class for sending email.
type StatefulEmail struct {
	destination string
	body        string
}

var emailPool = &sync.Pool{
	New: func() interface{} {
		return new(StatefulEmail)
	},
}

// GetInstanceEmail get the instance from the emailPool.
func GetInstanceEmail() *StatefulEmail {
	return emailPool.Get().(*StatefulEmail)
}

// PutInstanceEmail put the required instance back to the emailPool.
func PutInstanceEmail(instance *StatefulEmail) {
	emailPool.Put(instance)
}

// Reset all states to the default beginning.
func (se *StatefulEmail) Reset() {
	se.destination = ""
	se.body = ""
}

// SetTemplate set the destination and body of the email.
func (se *StatefulEmail) SetTemplate(dest, body string) {
	se.destination = dest
	se.body = body
}

// SendEmail sleep and print the destination and body variable of this class.
func (se *StatefulEmail) SendEmail() {
	// Doing sending email work
	time.Sleep(time.Second)
	fmt.Printf("Email send to: %s \nbody: %s \n", se.destination, se.body)
}
