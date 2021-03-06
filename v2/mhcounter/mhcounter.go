// Package mhcounter track the cumulative counts over the past minute and over the past hour.
// Useful, for example, to track recent bandwidth usage.
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
	var minuteEvents []event
	var hourEvents []event

	minuteAgo := now.Add(-1 * time.Minute)
	hourAgo := now.Add(-1 * time.Hour)

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

// Add a new data point (count >= 0).
// For the next minute, MinuteCount() will be larger by +count.
// For the next hour, HourCount() will be larger by +count.
func (e *Events) Add(count int) {
	var now = time.Now()
	e.shiftOldEvents(now)

	e.minuteEvents = append(e.minuteEvents, event{count, now})

	e.minuteCount += count
	e.hourCount += count
}

// MinuteCount returns the accumulated count over the past 60 seconds.
func (e *Events) MinuteCount() int {
	e.shiftOldEvents(time.Now())
	return e.minuteCount
}

// HourCount returns the accumulated count over the past 3600 seconds.
func (e *Events) HourCount() int {
	e.shiftOldEvents(time.Now())
	return e.hourCount
}
