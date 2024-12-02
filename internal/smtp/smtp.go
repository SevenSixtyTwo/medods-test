package smtp

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

type (
	SmtpBase struct {
		From     string
		To       string
		Password string
		Server   string
	}

	SmtpMail struct {
		Headers map[string]string
		From    *mail.Address
		To      *mail.Address
		Auth    *smtp.Auth
		Config  *tls.Config
	}
)

func SendMail(mail_body string, mail *SmtpMail, base *SmtpBase) error {
	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	mail_conn, err := tls.Dial("tcp", base.Server, mail.Config)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	defer mail_conn.Close()

	mail_client, err := smtp.NewClient(mail_conn, mail.Config.ServerName)
	if err != nil {
		return fmt.Errorf("new client: %s", err)
	}
	defer mail_client.Close()

	if err = mail_client.Auth(*mail.Auth); err != nil {
		return fmt.Errorf("auth: %s", err)
	}

	mail_message := ""
	for k, v := range mail.Headers {
		mail_message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	mail_message += "\r\n" + mail_body

	if err := mail_client.Mail(mail.From.Address); err != nil {
		return fmt.Errorf("mail: %s", err)
	}

	if err := mail_client.Rcpt(mail.To.Address); err != nil {
		return fmt.Errorf("rcpt: %s", err)
	}

	mail_writer, err := mail_client.Data()
	if err != nil {
		return fmt.Errorf("data: %s", err)
	}

	_, err = mail_writer.Write([]byte(mail_message))
	if err != nil {
		return fmt.Errorf("write: %s", err)
	}

	if err = mail_writer.Close(); err != nil {
		return fmt.Errorf("close: %s", err)
	}

	if err = mail_client.Quit(); err != nil {
		return fmt.Errorf("quit: %s", err)
	}

	return nil
}

func InitSmtp(smtp_cred *SmtpBase) (*SmtpMail, error) {
	var smtpMail *SmtpMail = &SmtpMail{}

	smtpMail.Headers = make(map[string]string)

	smtpMail.Headers["To"] = ""

	smtpMail.To = &mail.Address{"", smtp_cred.To}

	smtpMail.From = &mail.Address{"", smtp_cred.From}
	mail_subj := "новый IP адрес"

	smtpMail.Headers["From"] = smtpMail.From.String()
	smtpMail.Headers["Subject"] = mail_subj

	smtp_server := smtp_cred.Server

	smtp_host, _, err := net.SplitHostPort(smtp_server)
	if err != nil {
		return nil, fmt.Errorf("split host port: %s", err)
	}

	mail_auth := smtp.PlainAuth("", smtp_cred.From, smtp_cred.Password, smtp_host)

	smtpMail.Auth = &mail_auth

	tls_config := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtp_host,
	}

	smtpMail.Config = tls_config

	return smtpMail, nil
}
