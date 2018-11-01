package main

import (
	"fmt"
	"time"
)

func main() {
	var e Events
	e.Add(1)
	e.Add(2)
	e.Add(3)
	e.MinuteCount()
}

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
	fmt.Println(e.event)
	return len(e.event)
}

func (e *Events) HourCount() int {
	fmt.Println(e.event)
	return len(e.event)
}
