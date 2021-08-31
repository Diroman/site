package mail

import (
	"fmt"
	"net/smtp"

	"lawSite/internal/config"
)

type Client struct {
	email    string
	password string
	smtpHost string
	smtpPort string
	message  string
}

const linkTemplate = "https://lawbox-portal.ru/user/new-password/"

func NewClient(cfg *config.ConfigStruct) *Client {
	cfgMail := cfg.Mail
	return &Client{
		email:    cfgMail.Email,
		password: cfgMail.Password,
		smtpHost: cfgMail.Host,
		smtpPort: cfgMail.Port,
		message:  "Для сброса пароля перейдите по ссылке: \r\n",
	}
}

func (c *Client) SendPasswordUpdate(email, token string) error {
	auth := smtp.PlainAuth("", c.email, c.password, c.smtpHost)

	address := fmt.Sprintf("%s:%s", c.smtpHost, c.smtpPort)
	link := fmt.Sprintf("%s%s", linkTemplate, token)
	message := "Subject: Сброс пароля lawbox-portal.com\r\n" + "\r\n" +
		fmt.Sprintf("%s%s \r\n", c.message, link)
	err := smtp.SendMail(address, auth, c.email, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendMessage(email, message string) error {
	auth := smtp.PlainAuth("", c.email, c.password, c.smtpHost)
	address := fmt.Sprintf("%s:%s", c.smtpHost, c.smtpPort)

	message = "Subject: Feedback\r\n" + "\r\n" + message
	err := smtp.SendMail(address, auth, c.email, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
