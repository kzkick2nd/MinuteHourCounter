package trailingBucket

import (
	"time"

	"../conveyor"
)

type Counter struct {
	Buckets       *conveyor.QueneBuckets
	SecsPerBucket int
	updated       time.Time
}

func (c *Counter) update(now time.Time) {
	currentBucket := int(now.Unix() / int64(c.SecsPerBucket))
	lastUpdateBucket := int(c.updated.Unix() / int64(c.SecsPerBucket))
	c.Buckets.Shift(currentBucket - lastUpdateBucket)
	c.updated = now
}

func (c *Counter) Add(count int, now time.Time) {
	c.update(now)
	c.Buckets.AddToBack(count)
}

func (c *Counter) TrailingCount(now time.Time) int {
	c.update(now)
	return c.Buckets.TotalSum()
}
