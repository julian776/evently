package repositories

import (
	"context"
	"fmt"
	"net/smtp"
	"notifier/pkgs/emails/models"
)

func SendEmail(
	ctx context.Context,
	auth models.Auth,
	to []string,
	message []byte,
) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Create authentication
	smtpAuth := smtp.PlainAuth("", auth.Email, auth.Password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, smtpAuth, auth.Email, to, message)
	if err != nil {
		return fmt.Errorf("error sending email %s", err.Error())
	}
	return nil
}
