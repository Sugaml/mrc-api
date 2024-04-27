package main

import (
	"fmt"

	"github.com/Sugaml/mrc-api/api/utils/mailer"
)

func main() {
	sender := mailer.NewGmailSender("Babulal", "tamangsugam09@gmail.com", "eokbjhvwagftsaca")
	subject := "A test email"
	content := `
	<h1>Welcome to page</h1>
	`
	to := []string{"babulaltamamng@gmail.com"}
	attachFiles := []string{}
	err := sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	if err != nil {
		fmt.Print(fmt.Errorf("error to send email", err))
		return
	}
	fmt.Print("Success")
}
