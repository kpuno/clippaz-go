package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/k3a/html2text"
	"github.com/kpuno/clippaz/initializers"
	"github.com/kpuno/clippaz/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, data *EmailData) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	from := mail.NewEmail("clippaz", "clippaz.io@gmail.com")
	subject := data.Subject
	plainTextContent := "Clippaz verify email"
	to := mail.NewEmail(user.LastName+" "+user.FirstName, user.Email)

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verification-code.html", &data)

	htmlContent := html2text.HTML2Text(body.String())
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(config.SendGridApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Headers)
	}
}
