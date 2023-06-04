package ctxvalue

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type S struct {
	X int
}

func TestGetAndSet(t *testing.T) {
	key := Key()
	ctx := context.Background()

	v, ok := Get[*S](ctx, key)
	assert.Nil(t, v)
	assert.False(t, ok, "context value associated with key is not found")

	ctx = Set(ctx, key, &S{X: 100})

	v2, ok2 := Get[*S](ctx, key)
	assert.Equal(t, &S{X: 100}, v2, "Get returns value associated with key")
	assert.True(t, ok2)
}

func ExampleGet() {
	key := Key()
	ctx := Set(context.Background(), key, &S{X: 100})
	v, ok := Get[*S](ctx, key)
	fmt.Printf("%#v\n", v)
	fmt.Println(ok)

	// Output:
	// &ctxvalue.S{X:100}
	// true
}
