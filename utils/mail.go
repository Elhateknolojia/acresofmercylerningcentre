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
        "mail.acresofmercylearningcentre.sc.ke", // SMTP host
        587,                                    // Port (STARTTLS)
        from,                                   // Username (full email)
        os.Getenv(passEnv),                     // Password from env
    )

    // STARTTLS requires TLS config but not full SSL
    d.TLSConfig = &tls.Config{
        InsecureSkipVerify: true,
        ServerName: "mail.acresofmercylearningcentre.sc.ke",
    }

    return d.DialAndSend(m)
}
