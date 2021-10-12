package helpers

import (
	"math"
	"time"
)

// IsDeletable check if the tweet is within the offset user provided. (offset is in days)
func IsDeletable(createdAt time.Time, offset float64) bool {
	now := time.Now()
	t := math.Ceil(now.Sub(createdAt).Hours() / 24)
	if t > offset {
		return true
	}
	return false
}
