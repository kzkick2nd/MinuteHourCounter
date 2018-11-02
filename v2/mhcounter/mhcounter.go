package mhcounter

import "time"

type event struct {
	count int
	time  time.Time
}

type Events struct {
	minute_events []event
	hour_events   []event

	minute_count int
	hour_count   int
}

func (e *Events) Add(count int) {}

func (e *Events) MinuteCount() int {
	return e.minute_count
}

func (e *Events) HourCount() int {
	return e.hour_count
}
