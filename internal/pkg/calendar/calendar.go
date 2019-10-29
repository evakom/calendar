/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package calendar implements simple event calendar via protobuf.
package calendar

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"log"
)

// PrintTestData print test calendar scenarios
func PrintTestData(events interfaces.DB) {

	event1 := models.NewEvent()
	//out, err := proto.Marshal(event1)
	//if err != nil {
	//	log.Fatalln("Failed to encode event:", err)
	//}
	//
	//event1 = &Event{}
	//if err := proto.Unmarshal(out, event1); err != nil {
	//	log.Fatalln("Failed to parse event:", err)
	//}
	event1.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	_ = events.AddEvent(event1)

	event2 := models.NewEvent()
	event2.Subject = "222222222222222222222"
	event2.Body = "3333333333333333333"
	_ = events.AddEvent(event2)

	fmt.Printf("%+v\n", events.GetAllEvents())
	fmt.Println("Added event ^^^ --------------------------")

	if err := events.DelEvent(event1.ID); err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", events.GetAllEvents())
	fmt.Println("Deleted event ^^^ --------------------------")

	event2.User.Email = []string{"zzzzzzzzzzzzzzzz", "xxxxxxxxxxxxxxxxx"}
	if err := events.EditEvent(event2); err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", events.GetAllEvents())
	fmt.Println("Edit event ^^^ --------------------------")

	e2, err := events.GetOneEvent(event2.ID)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", e2)
	fmt.Println("Get one event ^^^ --------------------------")
}
