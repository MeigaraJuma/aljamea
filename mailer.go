package main

import (
	"crypto/tls"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func sendMail(email string, imams []string) {

	body := "<strong> Alhamdulillah Sudara Kita Yang Beruntung Jadi Imam </strong> \n"

	for i, user := range imams {
		body += fmt.Sprintf("%d. %s \n", i+1, user)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "nonatheist@noreply.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Yang Beruntung Jadi Imam")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "slimyhebat@gmail.com", os.Getenv("EMAIL_PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
