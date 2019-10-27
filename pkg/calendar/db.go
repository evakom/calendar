/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

// Constants
const (
	MapDBType = "map"
	//PostgresDBType = "postgres"
)

// DB is thw main interface for any DBs
type DB interface {
	AddEvent(event Event) error
	EditEvent(event Event) error
	DelEvent(id uint32) error
	GetEvent(id uint32) (Event, error)
	GetAllEvents() []Event
}

// NewDB returns DB by db type
func NewDB(dbType string) DB {
	switch dbType {
	case MapDBType:
		return newMapDB()
		//case PostgresDBType:
		//return newPostgresDB
	}
	return nil
}
