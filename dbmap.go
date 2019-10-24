/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/golang/protobuf/ptypes"
	"sync"
)

type dbMapEvents struct {
	sync.RWMutex
	events map[uint32]*Event
}

func (db *dbMapEvents) newDB() interface{} {
	return &dbMapEvents{
		events: make(map[uint32]*Event),
	}
}

func (db *dbMapEvents) addEvent(event *Event) error {
	db.Lock()
	defer db.Unlock()
	db.events[event.Id] = event
	return nil
}

func (db *dbMapEvents) delEvent(id uint32) error {
	db.Lock()
	defer db.Unlock()
	db.events[id].DeletedAt = ptypes.TimestampNow()
	return nil
}

func (db *dbMapEvents) editEvent(event *Event) error {
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = ptypes.TimestampNow()
	db.events[event.Id] = event
	return nil
}

func (db *dbMapEvents) getEvent(id uint32) *Event {
	return db.events[id]
}

func (db *dbMapEvents) getAllEvents() []*Event {
	events := make([]*Event, 0)
	for _, event := range db.events {
		events = append(events, event)
	}
	return events
}
