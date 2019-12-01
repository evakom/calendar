/*
 * HomeWork-14: RabbitMQ receiver
 * Created on 30.11.2019 23:30
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

// Constants.
const EmailTemplate = `
From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

// EmailMessage base struct.
type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type emailCredentials struct {
	userName, password string
	server             string
	port               int
}

var t *template.Template

func init() {
	t = template.New("email")
	if _, err := t.Parse(EmailTemplate); err != nil {
		log.Println(err)
	}
}

func sendEmail(from, subject, body string, to []string) error {
	message := &EmailMessage{
		From:    from,
		Subject: subject,
		Body:    body,
		To:      to,
	}

	var bodyB bytes.Buffer
	if err := t.Execute(&bodyB, message); err != nil {
		return err
	}

	authCreds := &emailCredentials{
		userName: "info",
		password: "",
		server:   "192.168.137.2",
		port:     25,
	}

	//auth := unencryptedAuth{smtp.PlainAuth("",
	//	authCreds.userName,
	//	authCreds.password,
	//	authCreds.server,
	//),
	//}
	auth := smtp.PlainAuth("",
		authCreds.userName,
		authCreds.password,
		authCreds.server,
	)

	sp := fmt.Sprintf("%s:%d", authCreds.server, authCreds.port)
	if err := smtp.SendMail(sp,
		auth,
		message.From,
		message.To,
		bodyB.Bytes(),
	); err != nil {
		return err
	}

	return nil
}

//type unencryptedAuth struct {
//	smtp.Auth
//}
//
//func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
//	s := *server
//	s.TLS = true
//	return a.Auth.Start(&s)
//}
