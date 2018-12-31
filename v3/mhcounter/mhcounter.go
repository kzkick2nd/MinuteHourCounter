// Package mhcounter track the cumulative counts over the past minute and over the past hour.
// Useful, for example, to track recent bandwidth usage.
package mhcounter

import (
	"time"

	"../conveyor"
	"../trailingBucket"
)

type Counter struct {
	minuteCounts trailingBucket.Counter
	hourCounts   trailingBucket.Counter
}

func NewCounter() *Counter {
	return &Counter{
		minuteCounts: trailingBucket.Counter{
			Buckets:       conveyor.NewBuckets(60),
			SecsPerBucket: 1,
		},
		hourCounts: trailingBucket.Counter{
			Buckets:       conveyor.NewBuckets(60),
			SecsPerBucket: 60,
		},
	}
}

// Add a new data point (count >= 0).
// For the next minute, MinuteCount() will be larger by +count.
// For the next hour, HourCount() will be larger by +count.
func (c *Counter) Add(count int) {
	var now = time.Now()
	c.minuteCounts.Add(count, now)
	c.hourCounts.Add(count, now)
}

// MinuteCount returns the accumulated count over the past 60 seconds.
func (c *Counter) MinuteCount() int {
	var now = time.Now()
	return c.minuteCounts.TrailingCount(now)
}

// HourCount returns the accumulated count over the past 3600 seconds.
func (c *Counter) HourCount() int {
	var now = time.Now()
	return c.hourCounts.TrailingCount(now)
}
