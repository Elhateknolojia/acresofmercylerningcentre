package utils

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
    "os"
)

func SendMail(from string, subject string, body string, passEnv string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", from)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(
        "mail.acresofmercylearningcentre.co.ke", // SMTP host
        465,                                    // Port (SSL)
        from,                                   // Username (full email)
        os.Getenv(passEnv),                     // Password from env
    )
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    return d.DialAndSend(m)
}
