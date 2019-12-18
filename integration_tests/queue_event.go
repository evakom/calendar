/*
 * HomeWork-9: Integration tests
 * Created on 18.12.2019 13:27
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/evakom/calendar/tools"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

const eventsQueueName = "events"

type alertTest struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	stop chan bool
	sync.RWMutex
	messages [][]byte
}

func (t *alertTest) startConsume(interface{}) {
	var err error
	t.messages = make([][]byte, 0)
	t.stop = make(chan bool)

	conf := tools.InitConfig("config.yml")

	t.conn, err = amqp.Dial(conf.RabbitMQ)
	if err != nil {
		log.Fatal(err)
	}

	t.ch, err = t.conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := t.ch.QueueDeclare(eventsQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	events, err := t.ch.Consume(q.Name, "", true, true, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func(stop <-chan bool) {
		for {
			select {
			case <-stop:
				return
			case event := <-events:
				t.Lock()
				t.messages = append(t.messages, event.Body)
				t.Unlock()
			}
		}
	}(t.stop)
}

func (t *alertTest) stopConsume(interface{}, error) {
	t.stop <- true

	errCh := t.ch.Close()
	errConn := t.conn.Close()

	if errCh != nil || errConn != nil {
		log.Println(errCh, errConn)
	}

	t.messages = nil
}

func iCreateEventWithEventRequestToServiceAPIWithOccursAtNow() error {
	return godog.ErrPending
}

func addedEventWillBeScheduledIntoMessageQueue() error {
	return godog.ErrPending
}

func getErrorHasNoError() error {
	return godog.ErrPending
}

func iConsumeMessageQueue() error {
	return godog.ErrPending
}

func iGetEventWithCorrectTestUserId() error {
	return godog.ErrPending
}

func willBeReadyToSendMessageRotThisUser() error {
	return godog.ErrPending
}

func FeatureContextQueueEvent(s *godog.Suite) {
	test := new(alertTest)
	s.BeforeScenario(test.startConsume)

	s.Step(`^I CreateEvent with EventRequest to service API with OccursAt = Now$`, iCreateEventWithEventRequestToServiceAPIWithOccursAtNow)
	s.Step(`^added event will be scheduled into message queue$`, addedEventWillBeScheduledIntoMessageQueue)
	s.Step(`^GetError has no error$`, getErrorHasNoError)
	s.Step(`^I consume message queue$`, iConsumeMessageQueue)
	s.Step(`^I get event with correct test user id$`, iGetEventWithCorrectTestUserId)
	s.Step(`^will be ready to send message rot this user$`, willBeReadyToSendMessageRotThisUser)

	s.AfterScenario(test.stopConsume)
}
