package multiton

import (
	"github.com/fairyhunter13/singleton-multiton/case1/multiton"
)

func sendEmail(destination, body string) {
	instanceEmail := multiton.GetInstanceEmail()
	defer func() {
		instanceEmail.Reset()
		multiton.PutInstanceEmail(instanceEmail)
	}()
	instanceEmail.SetTemplate(destination, body)
	instanceEmail.SendEmail()
}
