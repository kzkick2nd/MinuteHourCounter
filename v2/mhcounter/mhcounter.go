package mhcounter

import (
	"time"
)

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

func (e *Events) shiftOldEvents(now time.Time) {
	var minute_ago = now.Add(-1 * time.Minute)
	var hour_ago = now.Add(-1 * time.Hour)

	// Goでスライス減らすのつらいので、書籍のpopではなく新規sliceを用意（どうなん）
	var minute_events []event
	var hour_events []event

	for _, v := range e.minute_events {
		if v.time.Before(minute_ago) {
			e.hour_events = append(e.hour_events, v)
			e.minute_count -= v.count
		} else {
			minute_events = append(minute_events, v)
		}
	}
	e.minute_events = minute_events

	for _, v := range e.hour_events {
		if v.time.Before(hour_ago) {
			e.hour_count -= v.count
		} else {
			hour_events = append(hour_events, v)
		}
	}
	e.hour_events = hour_events
}

func (e *Events) Add(count int) {
	var now = time.Now()
	e.shiftOldEvents(now)

	e.minute_events = append(e.minute_events, event{count, now})

	e.minute_count += count
	e.hour_count += count
}

func (e *Events) MinuteCount() int {
	e.shiftOldEvents(time.Now())
	return e.minute_count
}

func (e *Events) HourCount() int {
	e.shiftOldEvents(time.Now())
	return e.hour_count
}
