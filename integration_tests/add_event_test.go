/*
 * HomeWork-9: Integration tests
 * Created on 16.12.2019 10:44
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import "github.com/DATA-DOG/godog"

func iSendCreateEventWithEventRequestToServiceAPI() error {
	return godog.ErrPending
}

func addedEventWillBeReturnedByGetEventWithIdOfTheEvent() error {
	return godog.ErrPending
}

func getErrorHasNoErrorsInBothCases() error {
	return godog.ErrPending
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
	s.Step(`^I send CreateEvent with EventRequest to service API$`, iSendCreateEventWithEventRequestToServiceAPI)
	s.Step(`^added event will be returned by GetEvent with id of the event$`, addedEventWillBeReturnedByGetEventWithIdOfTheEvent)
	s.Step(`^GetError has no errors in both cases$`, getErrorHasNoErrorsInBothCases)
	s.Step(`^I send GetEvent request with event id to service API$`, iSendGetEventRequestWithEventIdToServiceAPI)
	s.Step(`^I get EventResponse with id of the event$`, iGetEventResponseWithIdOfTheEvent)
	s.Step(`^GetError has no errors$`, getErrorHasNoErrors)
}
