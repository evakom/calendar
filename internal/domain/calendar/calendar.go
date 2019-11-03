/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package calendar implements simple event calendar via protobuf.
package calendar

import (
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
)

// Calendar is the main calendar struct.
type Calendar struct {
	db     interfaces.DB
	logger models.Logger
}

// NewCalendar inits main calendar fields.
func NewCalendar(db interfaces.DB) Calendar {
	return Calendar{
		db:     db,
		logger: models.Logger{}.GetLogger(),
	}
}

// AddEvent adds new event for given user.
func (c Calendar) AddEvent(event models.Event) error {
	c.logger.WithFields(models.Fields{
		"id": event.ID.String(),
	}).Info("Request add event into calendar")
	c.logger.Debug("Requested event body for adding into calendar: %+v", event)
	return c.db.AddEventDB(event)
}

// GetAllEvents returns all calendar events for given filter.
//func (c Calendar) GetAllEventsFilter(filter models.Event) ([]models.Event, error) {
//	result := c.db.GetAllEventsFilterDB(filter)
//
//	return result, nil
//}

// GetAllEventsByUserID returns all calendar events for given user.
func (c Calendar) GetAllEventsByUserID(userID string) ([]models.Event, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.ErrBadUserID
	}
	c.logger.Info("Requested events for user: ", userID)
	return c.GetAllEventsFilterDB(models.Event{UserID: uid})
}
