package mhcounter

import (
	"time"

	"../trailingBucket"
)

type Counter struct {
	minuteCounts trailingBucket.Counter
	hourCounts   trailingBucket.Counter
}

func New() Counter {
	return c.Counter{
		minuteCounts: trailingBucket.Counter{buckets: 60, secs: 1},
		hourCounts:   trailingBucket.Counter{buckets: 60, secs: 60},
	}
}

func (c *Counter) Add(count int) {
	var now = time.Now()
	e.minuteCounts.Add(count, now)
	e.hourCounts.Add(count, now)
}

func (c *Counter) MinuteCount() int {
	var now = time.Now()
	return c.minuteCounts.TrailingCount(now)
}

func (c *Counter) HourCount() int {
	var now = time.Now()
	return c.hourCounts.TrailingCount(now)
}
