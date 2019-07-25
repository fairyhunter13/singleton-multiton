package main

import (
	"github.com/fairyhunter13/singleton-multiton/case1/singleton"
)

var (
	singletonEmail = new(singleton.Email)
	instanceEmail  = singletonEmail.GetInstanceEmail()
)

func main() {
	instanceEmail.SetTemplate("test@example.com", "Hi! And Hello! test!")
	instanceEmail.SendEmail()
}
