package mailer

import (
	"crypto/tls"
	"fmt"

	"github.com/PureMLHQ/PureML/packages/purebackend/core/settings"
	gomail "gopkg.in/mail.v2"
)

func SendMail(settings *settings.Settings, to string, subject string, body string) (err error) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", settings.MailService.Username)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer(settings.MailService.Host, settings.MailService.Port, settings.MailService.Username, settings.MailService.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
