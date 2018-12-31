// Package mhcounter track the cumulative counts over the past minute and over the past hour.
// Useful, for example, to track recent bandwidth usage.
package mhcounter

import "time"

type event struct {
	count int
	time  time.Time
}

type Events struct {
	event []event
}

func (e *Events) countSince(cutoff time.Time) int {
	var count int
	for _, v := range e.event {
		if v.time.After(cutoff) {
			count += v.count
		}
	}
	return count
}

// Add a new data point (count >= 0).
// For the next minute, MinuteCount() will be larger by +count.
// For the next hour, HourCount() will be larger by +count.
func (e *Events) Add(count int) {
	e.event = append(e.event, event{count, time.Now()})
}

// MinuteCount returns the accumulated count over the past 60 seconds.
func (e *Events) MinuteCount() int {
	return e.countSince(time.Now().Add(-1 * time.Minute))
}

// HourCount returns the accumulated count over the past 3600 seconds.
func (e *Events) HourCount() int {
	return e.countSince(time.Now().Add(-1 * time.Hour))
}
