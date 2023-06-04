package ctxvalue

import "context"

// Get returns the value associated with this context for key, or zero-value
// if no value is associated with key.
func Get[T any](ctx context.Context, key *ContextKey) (val T, ok bool) {
	if v, ok := ctx.Value(key).(T); ok {
		return v, true
	}

	return *new(T), false
}

// Set returns a copy of parent in which the value associated with key is val.
func Set[T any](ctx context.Context, key *ContextKey, val T) context.Context {
	return context.WithValue(ctx, key, val)
}
