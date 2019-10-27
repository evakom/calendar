/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

type db interface {
	newDB() interface{}
	AddEvent(event Event) error
	EditEvent(event Event) error
	DelEvent(id uint32) error
	GetEvent(id uint32) (Event, error)
	GetAllEvents() []Event
}

// NewDB returns new db interface.
func NewDB(d db) interface{} {
	return d.newDB()
}
