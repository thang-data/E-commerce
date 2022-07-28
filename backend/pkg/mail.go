package pkg

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/thang-data/backend/config"
	"html/template"
	"net/smtp"
)

func bindingDataToTemplate(pathTemplate string, dataObject map[string]interface{}) (bytes.Buffer, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles(pathTemplate)
	if err != nil {
		return body, err
	}
	err = t.Execute(&body, dataObject)
	return body, err
}

// BindingDataToTemplateMail bind to local template
func BindingDataToTemplateMail(nameTemplate string, dataObject map[string]interface{}) (bytes.Buffer, error) {
	pathTemplate := fmt.Sprintf("assets/email_templates/%s", nameTemplate)
	return bindingDataToTemplate(pathTemplate, dataObject)
}

// BindingDataToTemplateMailInDB bind to template in DB
func BindingDataToTemplateMailInDB(templateHtml string, dataObject map[string]interface{}) (bytes.Buffer, error) {
	var body bytes.Buffer
	template.New("html").Parse(templateHtml)
	t, err := template.New("html").Parse(templateHtml)
	if err != nil {
		return body, err
	}
	err = t.Execute(&body, dataObject)
	return body, err
}

func SendEmailToMultipleAddress(to []string, subject string, body bytes.Buffer) error {
	cfg := config.GetConfig()

	if len(cfg.SendGridApiKey) > 0 { // send by SendGrid
		for i := range to {
			sgFrom := mail.NewEmail(cfg.SendGridSenderName, cfg.SendGridSenderAddress)
			sgTo := mail.NewEmail(to[i], to[i])
			message := mail.NewSingleEmail(sgFrom, subject, sgTo, "", body.String())
			client := sendgrid.NewSendClient(cfg.SendGridApiKey)
			_, err := client.Send(message)

			if err != nil {
				return err
			}
		}
	} else { // send by SMTP
		e := email.NewEmail()
		e.From = fmt.Sprintf("%s<%s> ", cfg.SMTPDisplayName, cfg.SMTPUsername)
		e.Subject = subject
		e.To = to

		e.HTML = body.Bytes()
		err := e.Send(fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort),
			smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPHost))
		if err != nil {
			return err
		}
	}

	return nil
}
