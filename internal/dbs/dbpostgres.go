/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 16:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"github.com/jmoiron/sqlx"
)

// TODO into config
const dsn = ""

// DBMapEvents is the base struct for using map db.
type DBPostgresEvents struct {
	db  *sqlx.DB
	err error
}

// NewPostgresDB returns new postgres db struct.
func NewPostgresDB() *DBPostgresEvents {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return &DBPostgresEvents{err: err}
	}
	err = db.Ping()
	if err != nil {
		return &DBPostgresEvents{err: err}
	}
	return &DBPostgresEvents{db: db}
}

//// AddEvent adds event to map db.
//func (db *DBMapEvents) AddEvent(event models.Event) error {
//	db.Lock()
//	defer db.Unlock()
//	db.Events[event.ID] = event
//	return nil
//}
//
//// DelEvent deletes one event by id.
//func (db *DBMapEvents) DelEvent(id uuid.UUID) error {
//	if _, ok := db.Events[id]; !ok {
//		return errors.ErrEventNotFound
//		//return fmt.Errorf("event id = %d not found", id)
//	}
//	db.Lock()
//	defer db.Unlock()
//	e := db.Events[id]
//	e.DeletedAt = time.Now()
//	db.Events[id] = e
//	return nil
//}
//
//// EditEvent updates one event.
//func (db *DBMapEvents) EditEvent(event models.Event) error {
//	if _, ok := db.Events[event.ID]; !ok {
//		return errors.ErrEventNotFound
//		//return fmt.Errorf("event id = %d not found", event.ID)
//	}
//	db.Lock()
//	defer db.Unlock()
//	event.UpdatedAt = time.Now()
//	db.Events[event.ID] = event
//	return nil
//}
//
//// GetOneEvent returns one event by id.
//func (db *DBMapEvents) GetOneEvent(id uuid.UUID) (models.Event, error) {
//	if _, ok := db.Events[id]; !ok {
//		return models.Event{}, errors.ErrEventNotFound
//		//return Event{}, fmt.Errorf("event id = %d not found", id)
//	}
//	if !db.Events[id].DeletedAt.IsZero() {
//		return models.Event{}, errors.ErrEventAlreadyDeleted
//		//return Event{}, fmt.Errorf("event id = %d already deleted", id)
//	}
//	return db.Events[id], nil
//}
//
//// GetAllEvents return all events slice.
//func (db *DBMapEvents) GetAllEvents() []models.Event {
//	events := make([]models.Event, 0)
//	for _, event := range db.Events {
//		events = append(events, event)
//	}
//	return events
//}
