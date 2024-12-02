package smtp

import (
	"fmt"
	"medods-test/internal/env"
)

func SendWarning(email, ip string) error {
	smtpBase := &SmtpBase{
		From:     env.SMTP_FROM,
		To:       email,
		Password: env.SMTP_PASSWORD,
		Server:   env.SMTP_SERVER,
	}

	smtpMial, err := InitSmtp(smtpBase)
	if err != nil {
		return fmt.Errorf("init smtp: %s", err)
	}

	mailBody := fmt.Sprintf("обнаружена попытка обновления токенов на ваш аккаунт с нового ip адреса\r\n ip адрес: %s", ip)

	if err := SendMail(mailBody, smtpMial, smtpBase); err != nil {
		return fmt.Errorf("send mail: %s", err)
	}

	return nil
}
