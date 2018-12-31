package mhcounter

import "time"

type event struct {
	count int
	time  time.Time
}

type Events struct {
	minuteEvents []event
	hourEvents   []event

	minuteCount int
	hourCount   int
}

func (e *Events) shiftOldEvents(now time.Time) {
	var minuteAgo = now.Add(-1 * time.Minute)
	var hourAgo = now.Add(-1 * time.Hour)

	var minuteEvents []event
	var hourEvents []event

	for _, v := range e.minuteEvents {
		if v.time.Before(minuteAgo) {
			e.hourEvents = append(e.hourEvents, v)
			e.minuteCount -= v.count
		} else {
			minuteEvents = append(minuteEvents, v)
		}
	}
	e.minuteEvents = minuteEvents

	for _, v := range e.hourEvents {
		if v.time.Before(hourAgo) {
			e.hourCount -= v.count
		} else {
			hourEvents = append(hourEvents, v)
		}
	}
	e.hourEvents = hourEvents
}

func (e *Events) Add(count int) {
	var now = time.Now()
	e.shiftOldEvents(now)

	e.minuteEvents = append(e.minuteEvents, event{count, now})

	e.minuteCount += count
	e.hourCount += count
}

func (e *Events) MinuteCount() int {
	e.shiftOldEvents(time.Now())
	return e.minuteCount
}

func (e *Events) HourCount() int {
	e.shiftOldEvents(time.Now())
	return e.hourCount
}
