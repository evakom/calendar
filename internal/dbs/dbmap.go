/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// DBMapEvents is the base struct for using map db.
type DBMapEvents struct {
	sync.RWMutex
	events map[uuid.UUID]models.Event
}

// NewMapDB returns new map db struct.
func NewMapDB() *DBMapEvents {
	return &DBMapEvents{
		events: make(map[uuid.UUID]models.Event),
	}
}

// AddEvent adds event to map db.
func (db *DBMapEvents) AddEvent(event models.Event) error {
	db.Lock()
	defer db.Unlock()
	db.events[event.ID] = event
	return nil
}

// DelEvent deletes one event by id.
func (db *DBMapEvents) DelEvent(id uuid.UUID) error {
	if _, ok := db.events[id]; !ok {
		return fmt.Errorf("event id = %d not found", id)
	}
	db.Lock()
	defer db.Unlock()
	e := db.events[id]
	e.DeletedAt = time.Now()
	db.events[id] = e
	return nil
}

// EditEvent updates one event.
func (db *DBMapEvents) EditEvent(event models.Event) error {
	if _, ok := db.events[event.ID]; !ok {
		return fmt.Errorf("event id = %d not found", event.ID)
	}
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = time.Now()
	db.events[event.ID] = event
	return nil
}

// GetOneEvent returns one event by id.
func (db *DBMapEvents) GetOneEvent(id uuid.UUID) (models.Event, error) {
	if _, ok := db.events[id]; !ok {
		return models.Event{}, fmt.Errorf("event id = %d not found", id)
	}
	if !db.events[id].DeletedAt.IsZero() {
		return models.Event{}, fmt.Errorf("event id = %d already deleted", id)
	}
	return db.events[id], nil
}

// GetAllEvents return all events slice.
func (db *DBMapEvents) GetAllEvents() []models.Event {
	events := make([]models.Event, 0)
	for _, event := range db.events {
		events = append(events, event)
	}
	return events
}
