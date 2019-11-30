/*
 * HomeWork-14: RabbitMQ receiver
 * Created on 30.11.2019 23:30
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strconv"
)

const EmailTemplate = `
From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	UserName, Password, Server string
	Port                       int
}

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(EmailTemplate)
}

func send() {
	message := &EmailMessage{
		From:    "tirava@gmail.com",
		Subject: "Testik",
		Body:    "Hello femina!",
		To:      []string{"info@feminasodt.ru"},
	}

	var body bytes.Buffer
	t.Execute(&body, message)

	authCreds := &EmailCredentials{
		UserName: "info",
		Password: "",
		Server:   "192.168.137.2",
		Port:     25,
	}

	auth := smtp.PlainAuth("",
		authCreds.UserName,
		authCreds.Password,
		authCreds.Server)

	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		body.Bytes())
}
