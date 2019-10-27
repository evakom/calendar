/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 27.10.2019 12:32
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	calendar2 "github.com/evakom/calendar/pkg/calendar"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	// ----------------- test code - will be deleted
	events := calendar2.NewDB(&calendar2.DBMapEvents{}).(*calendar2.DBMapEvents)

	event1 := calendar2.NewEvent()

	out, err := proto.Marshal(event1)
	if err != nil {
		log.Fatalln("Failed to encode event:", err)
	}

	event1 = &calendar2.Event{}
	if err := proto.Unmarshal(out, event1); err != nil {
		log.Fatalln("Failed to parse event:", err)
	}

	event1.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	_ = events.AddEvent(*event1)

	event2 := calendar2.NewEvent()
	event2.Subject = "222222222222222222222"
	event2.Body = "3333333333333333333"
	_ = events.AddEvent(*event2)

	fmt.Println(events.GetAllEvents())
	fmt.Println("--------------------------")

	_ = events.DelEvent(event1.Id)
	fmt.Println(events.GetAllEvents())
	fmt.Println("--------------------------")

	event2.User.Email = []string{"zzzzzzzzzzzzzzzz", "xxxxxxxxxxxxxxxxx"}
	_ = events.EditEvent(*event2)
	fmt.Println(events.GetAllEvents())
	fmt.Println("--------------------------")

	fmt.Println(events.GetEvent(2))
	fmt.Println("--------------------------")
	// ----------------- test code - will be deleted
}
