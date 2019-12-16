/*
 * HomeWork-9: Integration tests
 * Created on 16.12.2019 10:44
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"errors"
	"github.com/DATA-DOG/godog"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

const (
	tsLayout = "2006-01-02 15:04:05"
)

type eventTest struct {
	req    *api.EventRequest
	resp   *api.EventResponse
	conn   *grpc.ClientConn
	client api.CalendarServiceClient
	ctx    context.Context
}

func (t *eventTest) start(interface{}) {
	var err error
	conf := tools.InitConfig("config.yml")
	t.conn, err = grpc.Dial(conf.ListenGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	t.client = api.NewCalendarServiceClient(t.conn)

	t.ctx = context.TODO()

	t.req = &api.EventRequest{
		OccursAt:   parseDateTime("2019-12-16 12:36:55", tsLayout),
		Subject:    "GoDog added event",
		Body:       "HomeWork-9: Integration tests",
		Location:   "Moscow",
		Duration:   parseDuration("1h"),
		UserID:     "a7fdcee4-8a27-4200-8529-c5336c886f77",
		AlertEvery: parseDuration("1m"),
	}
}

func (t *eventTest) stop(interface{}, error) {
	if err := t.conn.Close(); err != nil {
		log.Println(err)
	}
}

func (t *eventTest) iSendCreateEventWithEventRequestToServiceAPI() error {
	var err error

	t.resp, err = t.client.CreateEvent(t.ctx, t.req)
	if err != nil {
		return err
	}

	return nil
}

func (t *eventTest) addedEventWillBeReturnedByGetEventWithIdOfTheEvent() error {
	id := t.resp.GetEvent().Id
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	t.req.ID = id

	return nil
}

func (t *eventTest) getErrorHasNoErrorsInBothCases() error {
	respErr := t.resp.GetError()
	if respErr != "" {
		return errors.New(respErr)
	}

	return nil
}

func iSendGetEventRequestWithEventIdToServiceAPI() error {
	return godog.ErrPending
}

func iGetEventResponseWithIdOfTheEvent() error {
	return godog.ErrPending
}

func getErrorHasNoErrors() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	test := new(eventTest)
	s.BeforeScenario(test.start)

	s.Step(`^I send CreateEvent with EventRequest to service API$`,
		test.iSendCreateEventWithEventRequestToServiceAPI)
	s.Step(`^added event will be returned by GetEvent with id of the event$`,
		test.addedEventWillBeReturnedByGetEventWithIdOfTheEvent)
	s.Step(`^GetError has no errors in both cases$`,
		test.getErrorHasNoErrorsInBothCases)
	s.Step(`^I send GetEvent request with event id to service API$`, iSendGetEventRequestWithEventIdToServiceAPI)
	s.Step(`^I get EventResponse with id of the event$`, iGetEventResponseWithIdOfTheEvent)
	s.Step(`^GetError has no errors$`, getErrorHasNoErrors)

	s.AfterScenario(test.stop)
}
