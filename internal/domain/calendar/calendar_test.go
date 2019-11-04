/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 04.11.2019 18:02
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

import (
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/tools"
	"testing"
)

const fileConfigPath = "../../../config.yml"

var cal Calendar

func init() {
	conf := tools.InitConfig(fileConfigPath)
	models.NewLogger("none", nil)
	db := tools.InitDB(conf.DBType)
	cal = NewCalendar(db)
}

func TestAddEvent(t *testing.T) {
	e := models.NewEvent()
	e.Subject = "44444444444444444"
	e.Body = "55555555555555555"
	if err := cal.AddEvent(e); err != nil {
		t.Errorf("Adding calendar event should not return error but returns it: %s", err)
	}
	e.Duration = 666
	if err := cal.AddEvent(e); err == nil {
		t.Errorf("Adding event with same id should return error but returns no error")
	}
}
