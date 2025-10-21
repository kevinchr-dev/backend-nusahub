package idgen

import (
	"time"
)

// Generate returns unix time in milliseconds as uint64
func Generate() uint64 {
	return uint64(time.Now().UnixMilli())
}

// NowMillis converts the id (which is unix ms) to int64 milliseconds
func NowMillis(id uint64) int64 {
	return int64(id)
}
