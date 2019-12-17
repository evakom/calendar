/*
 * HomeWork-9: Integration tests
 * Created on 17.12.2019 17:24
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"errors"
	"github.com/DATA-DOG/godog"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/evakom/calendar/tools"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"time"
)

type eventsTest struct {
	req    *api.EventRequest
	resp   *api.EventResponse
	conn   *grpc.ClientConn
	client api.CalendarServiceClient
	ctx    context.Context
	lastID string
	failID string
}

func (t *eventsTest) start(interface{}) {
	var err error
	conf := tools.InitConfig("config.yml")
	t.conn, err = grpc.Dial(conf.ListenGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	t.client = api.NewCalendarServiceClient(t.conn)

	t.ctx = context.TODO()

	t.req = &api.EventRequest{
		Subject:    "GoDog list events",
		Body:       "HomeWork-9: Integration tests",
		Location:   "Moscow",
		Duration:   parseDuration("1h"),
		UserID:     "a7fdcee4-8a27-4200-8529-c5336c886f77",
		AlertEvery: parseDuration("1m"),
	}
}

func (t *eventsTest) stop(interface{}, error) {
	if err := t.conn.Close(); err != nil {
		log.Println(err)
	}
}

func (t *eventsTest) iSendCreateEventToServiceAPIForCycleWithEventsForSameUserAndStepDaysForOccursAt(
	numEvents, stepDays int) error {

	for i := 0; i < numEvents; i++ {
		deltaDays := time.Duration(stepDays*i) * time.Hour * 24
		occursAt, err := ptypes.TimestampProto(time.Now().Add(deltaDays))
		if err != nil {
			return err
		}
		t.req.OccursAt = occursAt
		t.resp, err = t.client.CreateEvent(t.ctx, t.req)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *eventsTest) allAddedEventsWillBeReturnedByGetUserEventsForGivenUser() error {
	return godog.ErrPending
}

func (t *eventsTest) getErrorHasNoErrorsInBothCases() error {
	respErr := t.resp.GetError()
	if respErr != "" {
		return errors.New(respErr)
	}

	return nil
}

func iSendGetEventsForDayRequestWithCurrentDayToServiceAPI() error {
	return godog.ErrPending
}

func iGetEventsResponseWithEventInItWithOccursAtInCurrentDay(arg1 int) error {
	return godog.ErrPending
}

func getErrorHasNoErrors() error {
	return godog.ErrPending
}

func iSendGetEventsForWeekRequestWithCurrentDayToServiceAPI() error {
	return godog.ErrPending
}

func iGetEventsResponseWithEventsInItWithOccursAtInNearWeek(arg1 int) error {
	return godog.ErrPending
}

func iSendGetEventsForMonthRequestWithCurrentDayToServiceAPI() error {
	return godog.ErrPending
}

func iGetEventsResponseWithEventsInItWithOccursAtInNearMonth(arg1 int) error {
	return godog.ErrPending
}

func FeatureContextListEvents(s *godog.Suite) {
	test := new(eventsTest)
	s.BeforeScenario(test.start)

	s.Step(`^I send CreateEvent to service API for cycle with (\d+) events for same user and step (\d+) days for OccursAt$`,
		test.iSendCreateEventToServiceAPIForCycleWithEventsForSameUserAndStepDaysForOccursAt)
	s.Step(`^all added events will be returned by GetUserEvents for given user$`,
		test.allAddedEventsWillBeReturnedByGetUserEventsForGivenUser)
	s.Step(`^GetError has no errors in both cases$`,
		test.getErrorHasNoErrorsInBothCases)
	s.Step(`^I send GetEventsForDay request with current day to service API$`, iSendGetEventsForDayRequestWithCurrentDayToServiceAPI)
	s.Step(`^I get EventsResponse with (\d+) event in it with OccursAt in current day$`, iGetEventsResponseWithEventInItWithOccursAtInCurrentDay)
	s.Step(`^GetError has no errors$`, getErrorHasNoErrors)
	s.Step(`^I send GetEventsForWeek request with current day to service API$`, iSendGetEventsForWeekRequestWithCurrentDayToServiceAPI)
	s.Step(`^I get EventsResponse with (\d+) events in it with OccursAt in near week$`, iGetEventsResponseWithEventsInItWithOccursAtInNearWeek)
	s.Step(`^I send GetEventsForMonth request with current day to service API$`, iSendGetEventsForMonthRequestWithCurrentDayToServiceAPI)
	s.Step(`^I get EventsResponse with (\d+) events in it with OccursAt in near month$`, iGetEventsResponseWithEventsInItWithOccursAtInNearMonth)

	s.AfterScenario(test.stop)
}
