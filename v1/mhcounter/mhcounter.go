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

func (e *Events) Add(count int) {
	e.event = append(e.event, event{count, time.Now()})
}

func (e *Events) MinuteCount() int {
	return e.countSince(time.Now().Add(-1 * time.Minute))
}

func (e *Events) HourCount() int {
	return e.countSince(time.Now().Add(-1 * time.Hour))
}
