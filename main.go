/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

//go:generate protoc --go_out=. calendar.proto

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	events := &Events{}
	events.addEvent(newEvent())

	out, err := proto.Marshal(events)
	if err != nil {
		log.Fatalln("Failed to encode event:", err)
	}

	events1 := &Events{}
	if err := proto.Unmarshal(out, events1); err != nil {
		log.Fatalln("Failed to parse event:", err)
	}

	event := newEvent()
	event.Subject = "222222222222222222222"
	event.Body = "3333333333333333333"
	events1.addEvent(event)

	fmt.Println(events1.GetEvents())
}
