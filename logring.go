// Package logring preserves a record of the most recent log entries.
package logring

import (
	"container/ring"
	"io"
	"os"
	"strings"
	"sync"
)

var count int
var countMutex sync.Mutex
var rl *ringLog
var rlMutex sync.Mutex

type ringLog struct {
	*ring.Ring
}

func (r *ringLog) Write(p []byte) (int, error) {
	rlMutex.Lock()
	rl.Ring.Value = string(p)
	rl.Ring = rl.Ring.Next()
	rlMutex.Unlock()
	countMutex.Lock()
	count++
	countMutex.Unlock()
	return len(p), nil
}

// Count reports the total number of messages logged.
func Count() int {
	return count
}

// Recent returns the most recently logged messages.
func Recent() []string {
	s := make([]string, 0, rl.Ring.Len())
	rl.Ring.Do(func(v interface{}) {
		if v != nil {
			s = append(s, strings.TrimSuffix(v.(string), "\n"))
		}
	})
	return s
}

// Writer returns an io.Writer suitable as an argument to log.SetOutput.
// It saves the the most recent log entries, while also writing to Stderr.
// The ringSize argument sets the maximum number of messages to keep.
func Writer(ringSize int) io.Writer {
	rl = &ringLog{ring.New(ringSize)}
	return io.MultiWriter(os.Stderr, rl)
}
