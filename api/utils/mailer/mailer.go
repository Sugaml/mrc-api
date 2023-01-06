package mailer

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mailjet/mailjet-apiv3-go/v3"
	"github.com/sirupsen/logrus"
)

type MailService interface {
	SendVerifyEmail(to string, token string) error
	SendResetPassword(to string, token string) error
}

type MailClient struct {
	mailClient *mailjet.Client
}

func NewMail() (*mailjet.Client, error) {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error getting env, not coming through %v", err)
	}
	logrus.Info("We are getting the env values")
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	return mailjetClient, nil
}
