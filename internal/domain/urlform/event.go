/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 17.11.2019 19:26
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package urlform implements www-url-from encode/decode of the models entities .
package urlform

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"time"
)

// Constants.
const (
	FormSubject  = "subject"
	FormBody     = "body"
	FormLocation = "location"
	FormDuration = "duration"
	FormUserID   = "user_id"
	FormEventID  = "event_id"
)

// Values is the base www-url-form values type.
type Values map[string]string

// DecodeID returns decoded string id to uuid.
func DecodeID(sid string) (uuid.UUID, error) {
	uid := tools.IDString2UUIDorNil(sid)
	if uid == uuid.Nil {
		return uid, fmt.Errorf("invalid id=%s", sid)
	}
	return uid, nil
}

// DecodeEvent returns decoded event from www-url-form values.
func (v Values) DecodeEvent() (models.Event, error) {

	duration, err := time.ParseDuration(v[FormDuration])
	if err != nil {
		return models.Event{}, err
	}

	userID, err := DecodeID(v[FormUserID])
	if err != nil {
		return models.Event{}, fmt.Errorf("illegal user id - %w", err)
	}

	event := models.NewEvent()

	event.Subject = v[FormSubject]
	event.Body = v[FormBody]
	event.Location = v[FormLocation]
	event.Duration = duration
	event.UserID = userID

	return event, nil
}
