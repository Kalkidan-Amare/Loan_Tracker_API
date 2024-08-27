package internal

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"regexp"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func SendEmail(to, subject, body string, smtpConfig SMTPConfig) error {
	from := smtpConfig.Username
	pass := smtpConfig.Password
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(smtpConfig.Host+":"+smtpConfig.Port,
		smtp.PlainAuth("", from, pass, smtpConfig.Host),
		from, []string{to}, []byte(msg))

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

func GenerateRandomToken(length int) (string, error) {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}

func ValidateEmail(email string) bool {
	// regular expression for validating an email
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}

//just simple validation
func ValidatePassword(password string) bool {
	return len(password) >= 8
}