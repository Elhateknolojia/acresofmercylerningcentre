package utils

import (
    "fmt"
    "net/smtp"
)

func SendMail(to string, subject string, body string) error {
    // Configure your SMTP server
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    sender := "your-email@gmail.com"       // replace with your Gmail
    password := "your-app-password"        // use Gmail App Password

    auth := smtp.PlainAuth("", sender, password, smtpHost)

    msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

    return smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{to}, msg)
}
