/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 21:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

import "testing"

func TestNewEvent(t *testing.T) {
	e1 := newEvent().GetId()
	e2 := newEvent().GetId()
	if e2 == e1 {
		t.Errorf("'id1 = %v' same as 'id2 = %v'", e1, e2)
	}
}

func TestNewDB(t *testing.T) {
	events := newDB(&dbMapEvents{})
	if _, ok := events.(*dbMapEvents); !ok {
		t.Errorf("Can't cast to *dbMapEvents type: %+v", events)
	}
}
