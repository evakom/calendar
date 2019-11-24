/*
 * HomeWork-12: gRPC client
 * Created on 24.11.2019 13:44
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"
)

const tsLayout = "2006-01-02 15:04:05"

func main() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewCalendarServiceClient(conn)

	occursAt, err := parseDateTime("2019-12-08 15:00:01")
	if err != nil {
		log.Fatal(err)
	}

	dura, err := parseDuration("1h")
	if err != nil {
		log.Fatal(err)
	}

	req := &api.EventRequest{
		OccursAt: occursAt,
		Subject:  "SSS",
		Body:     "BBB",
		Location: "LLL",
		Duration: dura,
		UserID:   "a7fdcee4-8a27-4200-8529-c5336c886f78",
	}

	resp, err := client.CreateEvent(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.GetError() != "" {
		log.Fatal(resp.GetError())
	}

	log.Println(resp.GetEvent())
}

func parseDateTime(s string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(tsLayout, s)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func parseDuration(s string) (*duration.Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return nil, err
	}
	dr := ptypes.DurationProto(d)
	return dr, nil
}
