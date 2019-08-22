package mail

import (
	"net/smtp"
	"os"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

var auth smtp.Auth
var mailId string
var mailPw string

func Auth() {
	mailId = os.Getenv("GMAIL_ID")
	mailPw = os.Getenv("GMAIL_PW")
	auth = smtp.PlainAuth("", mailId, mailPw, "smtp.gmail.com")
}

func New(to []string, subject string, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) Send() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, mailId, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}
