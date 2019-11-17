/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 16.11.2019 23:13
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package json implements json encode/decode of the models entities.
package json

import (
	"encoding/json"
	"github.com/evakom/calendar/internal/domain/models"
)

// EventResult model.
type EventResult struct {
	models.Event `json:"result"`
}

// NewEventResult returns result struct.
func NewEventResult(event models.Event) EventResult {
	return EventResult{event}
}

// Encode marshals event to json.
func (er EventResult) Encode() (string, error) {
	b, err := json.Marshal(er)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
