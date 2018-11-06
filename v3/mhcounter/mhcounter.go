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

func (c *Counter) Add(count int) {
	var now = time.Now()
	c.minuteCounts.Add(count, now)
	c.hourCounts.Add(count, now)
}

func (c *Counter) MinuteCount() int {
	var now = time.Now()
	return c.minuteCounts.TrailingCount(now)
}

func (c *Counter) HourCount() int {
	var now = time.Now()
	return c.hourCounts.TrailingCount(now)
}
