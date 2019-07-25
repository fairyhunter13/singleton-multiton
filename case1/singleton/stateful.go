package singleton

import (
	"fmt"
	"time"
)

// This file is in another package.

// StatefulEmail is a class for sending email.
type StatefulEmail struct {
	destination string
	body        string
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

// Email contains StatefulEmail class for sending the email.
type Email struct {
	singleton *StatefulEmail
}

// GetInstanceEmail get the singleton instance from this variable.
func (email *Email) GetInstanceEmail() *StatefulEmail {
	if email.singleton == nil {
		email.singleton = new(StatefulEmail)
	}
	return email.singleton
}
