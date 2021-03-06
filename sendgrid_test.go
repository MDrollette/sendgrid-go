package sendgrid

import (
	"net/mail"
	"os"
	"testing"
)

func Test_Send(t *testing.T) {
	sg := NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"))
	message := NewMail()
	message.AddTo("Yamil Asusta <yamil.asusta@sendgrid.com>")
	address, _ := mail.ParseAddress("Yamil Asusta <yamil.asusta@upr.edu>")
	message.AddReceipient(address)
	message.AddSubject("SendGrid Testing")
	message.AddHTML("WIN")
	message.AddFrom("yamil@sendgrid.com")
	message.AddSubstitution("key", "value")
	filepath, _ := os.Getwd()
	message.AddAttachment(filepath + "/sendgrid.go")
	if r := sg.Send(message); r == nil {
		t.Log("Test_Send Passed")
	} else {
		t.Error("Test_Send Failed", r)
	}
}
