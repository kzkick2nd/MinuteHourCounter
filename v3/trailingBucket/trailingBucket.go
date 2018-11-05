package trailingBucket

import (
	"time"

	"../conveyor"
)

type Counter struct {
	secs    int
	updated time.Time

	buckets conveyor.QueneBuckets
}
