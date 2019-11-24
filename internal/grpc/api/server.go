/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 21:50
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package api implements grpc api.
package api

import (
	"context"
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/loggers"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//go:generate protoc --go_out=plugins=grpc:. --proto_path=../../../api api.proto

// CalendarServer base struct.
type CalendarServer struct {
	calendar calendar.Calendar
	logger   loggers.Logger
}

// NewCalendarServer returns new server struct.
func NewCalendarServer(cal calendar.Calendar) *CalendarServer {
	return &CalendarServer{
		calendar: cal,
		logger:   loggers.GetLogger(),
	}
}

// StartGRPCServer is registers and runs the server.
func (cs *CalendarServer) StartGRPCServer(addr string) {

	srv := grpc.NewServer()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		cs.logger.Error(err.Error())
		os.Exit(1)
	}

	RegisterCalendarServiceServer(srv, cs)

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		cs.logger.Warn("Signal received: %s", <-shutdown)
		srv.GracefulStop()
	}()

	cs.logger.Info("Starting gRPC server at: %s", addr)
	if err := srv.Serve(lis); err != nil {
		cs.logger.Error(err.Error())
		os.Exit(1)
	}
	cs.logger.Info("Shutdown gRPC server at: %s", addr)
}

// CreateEvent creates event.
func (cs *CalendarServer) CreateEvent(context.Context, *Event) (*EventResponse, error) {
	panic("implement me")
}

// GetEvent got one event by id.
func (cs *CalendarServer) GetEvent(context.Context, *ID) (*EventResponse, error) {
	panic("implement me")
}

// GetUserEvents returns all events for given user.
func (cs *CalendarServer) GetUserEvents(context.Context, *ID) (*EventsResponse, error) {
	panic("implement me")
}

// UpdateEvent updates event by id.
func (cs *CalendarServer) UpdateEvent(context.Context, *Event) (*EventResponse, error) {
	panic("implement me")
}

// DeleteEvent deletes event from DB.
func (cs *CalendarServer) DeleteEvent(context.Context, *ID) (*EventResponse, error) {
	panic("implement me")
}

// GetEventsForDay returns all events for given day.
func (cs *CalendarServer) GetEventsForDay(context.Context, *Day) (*EventsResponse, error) {
	panic("implement me")
}

// GetEventsForWeek returns all events for given week from day.
func (cs *CalendarServer) GetEventsForWeek(context.Context, *Day) (*EventsResponse, error) {
	panic("implement me")
}

// GetEventsForMonth returns all events for given month from day.
func (cs *CalendarServer) GetEventsForMonth(context.Context, *Day) (*EventsResponse, error) {
	panic("implement me")
}
