package email

import (
	"log"
	"sync"

	"github.com/go-mail/mail"
	"github.com/rngallen/fiber-email/config"
)

type UserDetails struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type UserList struct {
	To      UserDetails `json:"to"`
	Subject string      `json:"subject"`
	Body    string      `json:"body"`
}

var Wg sync.WaitGroup

func SendNewsLetter(emailList []UserList) error {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered Error: %v\n", r)
		}
	}()

	host := config.Config("EMAIL_HOST.SERVER")
	port := config.IntConfig("EMAIL_HOST.PORT")
	username := config.Config("EMAIL_HOST.USER")
	password := config.Config("EMAIL_HOST.PASSWORD")

	d := mail.NewDialer(host, port, username, password)

	d.StartTLSPolicy = mail.MandatoryStartTLS

	sender, err := d.Dial()
	if err != nil {
		// log.Panicln(err)
		return err
	}

	email := mail.NewMessage()
	for index, user := range emailList {
		Wg.Add(1)
		email.SetHeader("From", username)
		email.SetAddressHeader("To", user.To.Address, user.To.Name)
		email.SetHeader("Subject", user.Subject)
		email.SetBody("text/html", user.Body)

		if err := mail.Send(sender, email); err != nil {
			log.Printf("%v Could not send eamil to %q : %v", index, user.To.Name, err)
			return err
		} else {
			log.Printf("%v Email sent to successfully to :%v \n", index, user.To.Name)
		}

	}
	return nil
}
