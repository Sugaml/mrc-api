package mailer

import (
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v3"
	"github.com/sirupsen/logrus"
)

func SendResetPassword(to string, token string) error {
	forgotUrl := os.Getenv("HOST_URL") + "/resetpassword/" + token
	mailjetClient, err := NewMail()
	if err != nil {
		return err
	}
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			Priority: 1,
			From: &mailjet.RecipientV31{
				Email: "tamangsugam09@gmail.com",
				Name:  "Sugam Lama",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  "User",
				},
			},
			Subject:  "Forget Password",
			TextPart: "Hi",
			HTMLPart: "<h3>Your forgot password link </h3>" + forgotUrl,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		logrus.Error("Error to send email ", err)
		return err
	}
	logrus.Info("reset password mail successfully sent.", res)
	return nil
}

func SendVerifyEmail(to string, token string) error {
	verifyUrl := os.Getenv("HOST_URL") + "/user/verify/" + token
	mailjetClient, err := NewMail()
	if err != nil {
		return err
	}
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			Priority: 1,
			From: &mailjet.RecipientV31{
				Email: "tamangsugam09@gmail.com",
				Name:  "Sugam Lama",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  "User",
				},
			},
			Subject:  "Email Verification ",
			TextPart: "Hello",
			HTMLPart: "<h3>Welcome to MRC. Please click the link to verify your email. </h3>" + verifyUrl,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		logrus.Error("Error to send email ", err)
		return err
	}
	logrus.Info("email verify mail successfully sent.", res)
	return nil
}
