package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type EmailData struct {
	From    string
	To      string
	Subject string
	Body    string
}

const Sender = "noreply@festa-notify.cf"
const CharSet = "UTF-8"

func Send(data EmailData) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		return err
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(data.To),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(data.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(data.Subject),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)

	fmt.Println(result)
	if err != nil {
		return err
	}

	return nil
}
