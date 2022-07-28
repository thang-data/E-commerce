package usecase

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thang-data/backend/config"
	"github.com/thang-data/backend/db"
	dbRepo "github.com/thang-data/backend/db/repository"
	"github.com/thang-data/backend/entity"
	"github.com/thang-data/backend/pkg"
	"time"
)

func Send02RegisterExistingEmailAddress(to string) error {
	title := "Registration confirmation"
	cfg := config.GetConfig()

	mapVal := map[string]interface{}{
		"TITLE":         title,
		"EMAIL_ADDRESS": to,
		"LOGIN_URL":     cfg.FrontendLoginUrl,
		"FORGET_URL":    cfg.FrontendForgotPasswordUrl,
		"YEAR":          time.Now().Year(),
	}

	var mailBody bytes.Buffer
	template, err := GetDefaultEmailTemplate(2)
	if err != nil {
		logrus.Error("Send02RegisterExistingEmailAddress get email template from DB error (will use local email template instead): " + err.Error())
		mailBody, err = pkg.BindingDataToTemplateMail("01_register.html", mapVal)
	} else {
		mailBody, err = pkg.BindingDataToTemplateMailInDB(template.Message, mapVal)
	}
	if err != nil {
		return err
	}

	err = pkg.SendEmailToMultipleAddress([]string{to}, title, mailBody)
	if err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("Send02RegisterExistingEmailAddress(%s)", to))
	return nil
}
func Send03RegisterInformation(to string) error {
	title := "Registration information"
	cfg := config.GetConfig()

	mapVal := map[string]interface{}{
		"TITLE":         title,
		"EMAIL_ADDRESS": to,
		"LOGIN_URL":     cfg.FrontendLoginUrl,
		"FORGET_URL":    cfg.FrontendForgotPasswordUrl,
		"YEAR":          time.Now().Year(),
	}

	var mailBody bytes.Buffer
	template, err := GetDefaultEmailTemplate(3)
	if err != nil {
		logrus.Error("Send03RegisterInformation get email template from DB error (will use local email template instead): " + err.Error())
		mailBody, err = pkg.BindingDataToTemplateMail("02_register.html", mapVal)
	} else {
		mailBody, err = pkg.BindingDataToTemplateMailInDB(template.Message, mapVal)
	}
	if err != nil {
		return err
	}

	err = pkg.SendEmailToMultipleAddress([]string{to}, title, mailBody)
	if err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("Send03RegisterExistingEmailAddress(%s)", to))
	return nil
}

// GetDefaultEmailTemplate without vars
func GetDefaultEmailTemplate(emailTypeNumber int) (*entity.EmailTemplate, error) {
	tx := db.Connect()
	emailTemplate, err := dbRepo.GetDefaultEmailTemplate(tx, getEmailTypeId(emailTypeNumber))
	if err != nil {
		return nil, err
	}

	return emailTemplate, nil
}

func getEmailTypeId(emailTypeNumber int) string {
	return fmt.Sprintf("00000000-0000-5661-7201-456d61696c%02d", emailTypeNumber)
}
