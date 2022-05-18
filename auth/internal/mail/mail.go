package mail

import (
	"errors"
	"fmt"
	"github.com/MoraGames/StreamingScheduler/auth/internal/utils"
	"net/smtp"
	"strings"
)

func SendEmail(serverAddress, mailAddress, mailKey, to, sub, tmpl string, data interface{}) error {
	from := mailAddress
	pass := mailKey

	emailBody, err := utils.ParseTemplate(tmpl, data)
	if err != nil {
		return errors.New("unable to parse email template")
	}

	fmt.Println(emailBody)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + sub + "\n\n" +
		emailBody

	err = smtp.SendMail(serverAddress,
		smtp.PlainAuth("", from, pass, strings.Split(serverAddress, ":")[0]),
		from, []string{to}, []byte(msg))

	if err != nil {
		return errors.New(fmt.Sprintf("smtp error: %s", err))
	}

	return nil
}
