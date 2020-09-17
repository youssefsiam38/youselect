package framework

import (
	"fmt"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendTokenEmail(token *string) error {
	link := fmt.Sprintf("%s/admin?token=%s", os.Getenv("FE_DOMAIN"), *token)
	to := mail.NewEmail("My Admin", os.Getenv("ADMIN_SECOND_EMAIL"))
	subject := "Your Token"
	from := mail.NewEmail("You Select", os.Getenv("ADMIN_EMAIL"))
	plainTextContent := fmt.Sprintf("Your link: %s", link)
	htmlContent := fmt.Sprintf("<a href=\"%s\" ><h1 style=\"text-align: center;\" >Press HERE</h1></a>", link)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}