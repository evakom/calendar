/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/evakom/calendar/tools"
	"io"
	"net/http"
)

type handler struct {
	handlers map[string]http.HandlerFunc
	calendar calendar.Calendar
	logger   loggers.Logger
	error    Error
}

func newHandlers(calendar calendar.Calendar) *handler {
	return &handler{
		handlers: make(map[string]http.HandlerFunc),
		calendar: calendar,
		logger:   loggers.GetLogger(),
		error:    newError(),
	}
}

func (h handler) hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "nobody"
	}
	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")

	s := "Hello, my name is " + name + "\n\n"

	if _, err := io.WriteString(w, s); err != nil {
		h.logger.Error("[hello] error write to response writer")
	}

	// test code for debug
	event := models.NewEvent()
	event.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	event.UserID = tools.IDString2UUIDorNil("a7fdcee4-8a27-4200-8529-c5336c886f77")
	_ = h.calendar.AddEvent(event)
}

func (h handler) getEvent(w http.ResponseWriter, r *http.Request) {
	key := "event_id"
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[getEvent] error: ", err)
	}
}

func (h handler) getUserEvents(w http.ResponseWriter, r *http.Request) {
	key := "user_id"
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[getUserEvents] error: ", err)
	}
}

func (h handler) createEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "create")
}

func (h handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update")
}

func (h handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "delete")
}

func (h handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for day")
}

func (h handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for week")
}

func (h handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for month")
}
