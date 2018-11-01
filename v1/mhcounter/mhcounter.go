package mhcounter

import "time"

type event struct {
	count int
	time  time.Time
}

type Events struct {
	event []event
}

func (e *Events) Add(count int) {
	e.event = append(e.event, event{count, time.Now()})
}

func (e *Events) MinuteCount() int {
	var count int
	var current = time.Now()
	for _, v := range e.event {
		if v.time.After(current.Add(-1 * time.Minute)) {
			count += v.count
		}
	}
	return count
}

func (e *Events) HourCount() int {
	var count int
	var current = time.Now()
	for _, v := range e.event {
		if v.time.After(current.Add(-1 * time.Hour)) {
			count += v.count
		}
	}
	return count
}
