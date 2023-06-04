package ctxvalue

import "sync/atomic"

// ContextKey represents a unique key for context value.
type ContextKey struct {
	id int64
}

var counter int64

// Key returns a new unique key of context value.
// Users should reuse the same key for the same purpose value.
func Key() *ContextKey {
	id := atomic.AddInt64(&counter, 1)

	return &ContextKey{
		id,
	}
}
