package emails

import (
	"context"
	"fmt"
	"net/smtp"
)

func SendEmail(
	ctx context.Context,
	auth Auth,
	to []string,
	message []byte,
) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	smtpAuth := smtp.PlainAuth("", auth.Email, auth.Password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, smtpAuth, auth.Email, to, message)
	if err != nil {
		return fmt.Errorf("error sending email %s", err.Error())
	}
	return nil
}
