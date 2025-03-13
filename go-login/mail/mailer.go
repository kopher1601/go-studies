package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mailer interface {
	SendWithActivateToken(email, token string) error
}

type mailhogMailer struct {
}

func NewMailhogMailer() Mailer {
	return &mailhogMailer{}
}

var (
	hostname = "mail"
	port     = 1025
	username = "user@example.com"
	password = "password"
)

func (m *mailhogMailer) SendWithActivateToken(email, token string) error {
	from := "info@login-go.app"
	recipients := []string{email}
	subject := "認証コード by LOGIN-GO"
	body := fmt.Sprintf("%s:%d", hostname, port)

	smtpServer := fmt.Sprintf("%s:%d", hostname, port)
	auth := smtp.CRAMMD5Auth(username, password)
	msg := []byte(strings.ReplaceAll(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, strings.Join(recipients, ","), subject, body), "\n", "\r\n"))
	if err := smtp.SendMail(smtpServer, auth, email, recipients, msg); err != nil {
		return err
	}
	return nil
}
