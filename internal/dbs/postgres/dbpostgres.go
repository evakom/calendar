/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 16:12
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package postgres implements postgres interface.
package postgres

import (
	"context"
	"fmt"
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib" // driver for postgres
	"github.com/jmoiron/sqlx"
	"time"
)

// Constants
const (
	EventIDField = "event_id"
	UserIDField  = "user_id"
	DayField     = "day"
	DeltaField   = "delta"
	EventsTable  = "events"
)

// DBPostgres is the base struct for using map db.
type DBPostgres struct {
	db     *sqlx.DB
	ctx    context.Context
	logger loggers.Logger
}

// CloseDB closes storage
func (db *DBPostgres) CloseDB() error {
	if err := db.db.Close(); err != nil {
		return err
	}
	db.logger.Info("Closed postgres DB")
	return nil
}

// NewPostgresDB returns new postgres db struct.
func NewPostgresDB(ctx context.Context, dsn string) (*DBPostgres, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error open db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error ping db: %w", err)
	}
	dbPg := &DBPostgres{
		db:     db,
		ctx:    ctx,
		logger: loggers.GetLogger(),
	}
	dbPg.logger.Info("Connected to postgres DB")
	return dbPg, nil
}

// AddEventDB adds event to postgres db.
func (db *DBPostgres) AddEventDB(event models.Event) error {
	query := "insert into " + EventsTable + " (id, createdat, updatedat, deletedat, occursat, " +
		"subject, body, duration, location, userid) " +
		"values(:id, :createdat, :updatedat, :deletedat, :occursat, " +
		":subject, :body, :duration, :location, :userid)"
	result, err := db.db.NamedExecContext(db.ctx, query, event)
	if err != nil {
		db.logger.Error("[AddEventDB][NamedExecContext]: %s", err)
		return errors.ErrEventAlreadyExists
	}

	ra, err := result.RowsAffected()
	if err != nil {
		db.logger.Error("[AddEventDB][RowsAffected]: %s", err)
		return fmt.Errorf("error get affected rows")
	}
	if ra != 1 {
		db.logger.Error("[AddEventDB][RowsAffected]: no affected")
		return fmt.Errorf("event not inserted into db: no rows affected")
	}

	db.logger.WithFields(loggers.Fields{
		EventIDField: event.ID.String(),
		UserIDField:  event.UserID.String(),
	}).Info("Event added into postgres DB")
	db.logger.Debug("Event body added: %+v", event)
	return nil
}

// DelEventDB deletes one event by id.
func (db *DBPostgres) DelEventDB(id uuid.UUID) error {
	event := models.Event{
		ID:        id,
		DeletedAt: time.Now(),
	}
	query := "update " + EventsTable + " set deletedat=:deletedat where id=:id"

	result, err := db.db.NamedExecContext(db.ctx, query, event)
	if err != nil {
		db.logger.Error("[DelEventDB][NamedExecContext]: %s", err)
		return fmt.Errorf("error execute delete event from DB")
	}

	ra, err := result.RowsAffected()
	if err != nil {
		db.logger.Error("[DelEventDB][RowsAffected]: %s", err)
		return fmt.Errorf("error get affected rows")
	}
	if ra == 0 {
		db.logger.Error("[DelEventDB][RowsAffected]: no affected")
		return errors.ErrEventNotFound
	}

	db.logger.WithFields(loggers.Fields{
		EventIDField: id.String(),
	}).Info("Event deleted from postgres DB")
	db.logger.Debug("Event body deleted from postgres DB: %+v", event)
	return nil
}

// EditEventDB updates one event.
func (db *DBPostgres) EditEventDB(event models.Event) error {
	eventNew := models.Event{
		ID:        event.ID,
		UpdatedAt: time.Now(),
		OccursAt:  event.OccursAt,
		Subject:   event.Subject,
		Body:      event.Body,
		Duration:  event.Duration,
		Location:  event.Location,
		UserID:    event.UserID,
	}
	query := "update " + EventsTable + " set updatedat=:updatedat, " +
		"occursat=:occursat, subject=:subject, body=:body, " +
		"duration=:duration, location=:location " +
		"where id=:id and deletedat =:deletedat"

	result, err := db.db.NamedExecContext(db.ctx, query, eventNew)
	if err != nil {
		db.logger.Error("[EditEventDB][NamedExecContext]: %s", err)
		return fmt.Errorf("error execute update event in DB")
	}

	ra, err := result.RowsAffected()
	if err != nil {
		db.logger.Error("[EditEventDB][RowsAffected]: %s", err)
		return fmt.Errorf("error get affected row")
	}
	if ra != 1 {
		db.logger.Error("[EditEventDB][RowsAffected]: no affected")
		return errors.ErrEventNotFound
		//return fmt.Errorf("event not updated in db: no rows affected")
	}

	db.logger.WithFields(loggers.Fields{
		EventIDField: eventNew.ID.String(),
		UserIDField:  eventNew.UserID.String(),
	}).Info("Event updated in postgres DB")
	db.logger.Debug("Event body updated in postgres DB: %+v", eventNew)
	return nil
}

// GetOneEventDB returns one event by id.
func (db *DBPostgres) GetOneEventDB(id uuid.UUID) (models.Event, error) {
	event := models.Event{ID: id}
	query := "select * from " + EventsTable + " where id=:id and deletedat =:deletedat"

	rows, err := db.db.NamedQueryContext(db.ctx, query, event)
	if err != nil {
		db.logger.Error("[GetOneEventDB][NamedQueryContext]: %s", err)
		return event, fmt.Errorf("error execute get one event from DB")
	}

	if rows.Next() {
		if err := rows.StructScan(&event); err != nil {
			db.logger.Error("[GetOneEventDB][StructScan]: %s", err)
			return event, fmt.Errorf("error scan DB row to event")
		}
	} else {
		return event, errors.ErrEventNotFound
	}

	db.logger.WithFields(loggers.Fields{
		EventIDField: id.String(),
		UserIDField:  event.UserID.String(),
	}).Info("Event got from postgres DB")
	db.logger.Debug("Event body got from postgres DB: %+v", event)

	if err := rows.Close(); err != nil {
		db.logger.Error("[GetOneEventDB] error close rows: %s", err)
	}
	return event, nil
}

// GetAllEventsDB return all events slice for given user id (no deleted).
func (db *DBPostgres) GetAllEventsDB(id uuid.UUID) []models.Event {
	events := make([]models.Event, 0)
	event := models.Event{UserID: id}
	query := "select * from " + EventsTable + " where userid=:userid and deletedat =:deletedat"

	rows, err := db.db.NamedQueryContext(db.ctx, query, event)
	if err != nil {
		db.logger.Error("[GetAllEventsDB][NamedQueryContext]: %s", err)
		return events
	}

	for rows.Next() {
		if err := rows.StructScan(&event); err != nil {
			db.logger.Error("[GetAllEventsDB][StructScan]: %s", err)
			continue
		}
		events = append(events, event)
	}

	db.logger.WithFields(loggers.Fields{
		UserIDField: id.String(),
	}).Info("All events [%d] got from postgres DB", len(events))

	if err := rows.Close(); err != nil {
		db.logger.Error("[GetAllEventsDB] error close rows: %s", err)
	}
	return events
}

// CleanEventsDB cleans db and deletes all events in the db for given user id (no restoring!).
func (db *DBPostgres) CleanEventsDB(id uuid.UUID) error {
	event := models.Event{UserID: id}

	uid := ""
	if id != uuid.Nil {
		uid = " where userid=:userid"
	}
	query := "delete from " + EventsTable + uid

	result, err := db.db.NamedExecContext(db.ctx, query, event)
	if err != nil {
		db.logger.Error("[CleanEventsDB][NamedExecContext]: %s", err)
		return fmt.Errorf("error execute delete events from DB: %w", err)
	}

	ra, err := result.RowsAffected()
	if err != nil {
		db.logger.Error("[CleanEventsDB][RowsAffected]: %s", err)
		return fmt.Errorf("error get affected rows: %w", err)
	}
	if ra == 0 {
		db.logger.Error("[CleanEventsDB][RowsAffected]: no affected")
		return errors.ErrEventsNotFound
	}

	db.logger.WithFields(loggers.Fields{
		UserIDField: id.String(),
	}).Info("All events [%d] deleted in postgres DB", ra)
	return nil
}

// GetAllEventsDBDays returns events for num of the days for given user
func (db *DBPostgres) GetAllEventsDBDays(filter models.Event) []models.Event {
	events := make([]models.Event, 0)
	event := models.Event{
		UserID:   filter.UserID,
		OccursAt: filter.OccursAt,
	}
	occursEnd := filter.OccursAt.Add(filter.Duration)
	event.UpdatedAt = occursEnd

	uid := ""
	if filter.UserID != uuid.Nil {
		uid = "userid=:userid and"
	}

	query := "select * from " + EventsTable + " where " + uid +
		" deletedat =:deletedat and occursat>=:occursat and occursat<:updatedat"

	rows, err := db.db.NamedQueryContext(db.ctx, query, event)
	if err != nil {
		db.logger.Error("[GetAllEventsDBDays][NamedQueryContext]: %s", err)
		return events
	}

	for rows.Next() {
		if err := rows.StructScan(&event); err != nil {
			db.logger.Error("[GetAllEventsDBDays][StructScan]: %s", err)
			continue
		}
		events = append(events, event)
	}

	db.logger.WithFields(loggers.Fields{
		DayField:   filter.OccursAt.String(),
		DeltaField: filter.Duration,
	}).Info("All events [%d] for day(s) got from map DB", len(events))
	return events
}
