package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, "traceID", traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value("traceID")

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}
