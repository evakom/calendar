/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:11
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package models implements base models.
package models

import (
	"github.com/google/uuid"
	"time"
)

const formatString = "%-13s%s\n"

// Event is the base event struct.
type Event struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	OccursAt  time.Time
	Subject   string
	Body      string
	Duration  time.Duration
	Location  string
	UserID    uuid.UUID
}

// NewEvent returns new example event.
func NewEvent() Event {
	return Event{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Subject:   "111",
		Body:      "222",
		Duration:  time.Minute,
		Location:  "Moscow",
		UserID:    uuid.New(), // todo uuid.Nil
	}
}
