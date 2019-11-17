/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/domain/json"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"io"
	"net/http"
)

const (
	eventIDField = "event_id"
	userIDField  = "user_id"
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
	//userID := query.Get("user_id")

	if name == "" {
		name = "nobody"
	}
	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")

	event := models.NewEvent()
	event.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	event.UserID = uuid.New()
	_ = h.calendar.AddEvent(event)

	s := "Hello, my name is " + name + "\n\n"

	if _, err := io.WriteString(w, s); err != nil {
		h.logger.Error("[hello] error write to response writer")
	}
}

func (h handler) getEvent(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	eventID := query.Get("event_id")
	events, err := h.calendar.GetAllEventsFilter(models.Event{
		ID: tools.IDString2UUIDorNil(eventID),
	})
	if err != nil {

		h.logger.WithFields(loggers.Fields{
			ReqIDField:   getRequestID(r.Context()),
			eventIDField: eventID,
		}).Error(err.Error())

		h.error.send(w, http.StatusOK, err, "error while get event id="+eventID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := json.NewEventResult(events).Encode()
	if err != nil {

		h.logger.WithFields(loggers.Fields{
			ReqIDField:   getRequestID(r.Context()),
			eventIDField: eventID,
		}).Error(err.Error())

		h.error.send(w, http.StatusOK, err, "error while encode event id="+eventID)
		return
	}
	if _, err := io.WriteString(w, result); err != nil {
		h.logger.Error("[getEvent] error write to response writer")
	}

	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")
}

func (h handler) getUserEvents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("user_id")
	events, err := h.calendar.GetAllEventsFilter(models.Event{
		UserID: tools.IDString2UUIDorNil(userID),
	})
	if err != nil {

		h.logger.WithFields(loggers.Fields{
			ReqIDField:  getRequestID(r.Context()),
			userIDField: userID,
		}).Error(err.Error())

		h.error.send(w, http.StatusOK, err, "error while get events for user id="+userID)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := json.NewEventResult(events).Encode()
	if err != nil {

		h.logger.WithFields(loggers.Fields{
			ReqIDField:   getRequestID(r.Context()),
			eventIDField: userID,
		}).Error(err.Error())

		h.error.send(w, http.StatusOK, err, "error while encode event id="+userID)
		return
	}
	if _, err := io.WriteString(w, result); err != nil {
		h.logger.Error("[getUserEvents] error write to response writer")
	}

	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")
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
