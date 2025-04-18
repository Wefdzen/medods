package service

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

type EmailService interface {
	SendWarningEmail(guid, oldIP, newIP string) error
}

func SendWarningEmail(es EmailService, guid, oldIp, newIp string) error {
	return es.SendWarningEmail(guid, oldIp, newIp)
}

// Реальная реализация (SMTP или другой сервис)
type RealEmailService struct{}

// SendWarningEmail отправляет сам себе плюс надо одобрить на этой почте api key
// работает с gmail you can change
func (s *RealEmailService) SendWarningEmail(guid, oldIP, newIP string) error {
	// Data for mail
	var (
		yourMail           = "yourEmail@gmail.com"
		passwordOfYourMail = "password"
		host               = "smtp.gmail.com"
	)

	// Set up authentication information.
	auth := smtp.PlainAuth("", yourMail, passwordOfYourMail, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{yourMail}

	msg := []byte(fmt.Sprintf("To: %v\r\n"+
		"Subject: medods Warning!\r\n"+
		"\r\n"+
		"was send refresh request to your account with not your ip you lose Your couple of tokens .\r\n", yourMail))
	// You can you yourself mail
	err := smtp.SendMail(fmt.Sprintf("%v:25", host), auth, yourMail, to, msg)
	if err != nil {
		return err
	}
	return nil
}

type StubEmailService struct{}

func (s *StubEmailService) SendWarningEmail(guid, oldIP, newIP string) error {
	log.Println("Типо отправилось сообщение если хотите ориг надо использвать RealEmailService а не Stub")
	log.Println("Guid: ", guid, "OldIP: ", oldIP, "newIP: ", newIP)
	return errors.New("NOT YOUR PC")
}
