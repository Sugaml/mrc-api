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

func SendStudentEnrollCompletedEmail(to string, name string) error {
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
			Subject:  " Enrollment Form Completed",
			TextPart: "Hello  " + name,
			HTMLPart: "<h3>I hope this email finds you well. I wanted to follow up with you regarding the enrollment form you recently completed. Thank you for taking the time to fill it out and provide us with the necessary information. </h3><br/><h3>Thank you again for your interest in our program, and we look forward to reviewing your application.</h3>",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		logrus.Error("Error to send email ", err)
		return err
	}
	logrus.Info("enrolled completed mail successfully sent.", res)
	return nil
}

func SendStudentEnrollApprovedEmail(to string, name string) error {
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
			Subject:  "Confirmation of Approved Enrollment Form",
			TextPart: "Hello  " + name,
			HTMLPart: "<h3>I am pleased to inform you that your enrollment form has been approved and you are now officially enrolled in BICT. Congratulations on taking this important step towards achieving your academic and professional goals!</h3><br/><h3>If you have any questions or concerns, please do not hesitate to reach out to us at support@mrc.com. We are here to support you throughout your academic journey.</h3><br/><h3>MRC</h3>",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		logrus.Error("Error to send email ", err)
		return err
	}
	logrus.Info("enrolled completed mail successfully sent.", res)
	return nil
}
