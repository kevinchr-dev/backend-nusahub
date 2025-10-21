package idgen

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// Generate generates a 64-bit ID that contains a timestamp (ms since epoch)
// in the high 42 bits and random data in the low 22 bits. This provides
// monotonic-ish ids with embedded time and a large random space.
func Generate() uint64 {
	// 42 bits for milliseconds timestamp (enough for many years)
	ts := uint64(time.Now().UnixNano()/int64(time.Millisecond)) & ((1 << 42) - 1)

	// 22 bits random
	var buf [8]byte
	_, _ = rand.Read(buf[:])
	rnd := binary.BigEndian.Uint64(buf[:]) & ((1 << 22) - 1)

	return (ts << 22) | rnd
}

// NowMillis extracts the millisecond timestamp portion from the ID
func NowMillis(id uint64) int64 {
	return int64(id >> 22)
}
