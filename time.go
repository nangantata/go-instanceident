package instanceident

import (
	"hash/adler32"
	"os"
	"time"
)

var lastTimeDerivedIdentitySequence int64 = 0

// TimeDerivedInt64Identity generates identity from hostname, PID and time.
// This function is NOT thread-safe.
// The clock resolution in identity generation is nanosecond.
func TimeDerivedInt64Identity() (ident int64) {
	hostname, err := os.Hostname()
	if nil != err {
		hostname = "-u"
	}
	hostHash := uint64(adler32.Checksum([]byte(hostname)))
	pidVal := (uint64(os.Getpid()) << 32) & 0x7FFFFFFF00000000
	clock := time.Now().UnixNano()
	if clock <= lastTimeDerivedIdentitySequence {
		clock = lastTimeDerivedIdentitySequence + 1
	}
	lastTimeDerivedIdentitySequence = clock
	ident = (int64)(((((hostHash << 32) ^ (hostHash << 48)) & 0x7FFF000000000000) ^ pidVal) | (uint64(clock) & 0xFFFFFFFF))
	return
}
