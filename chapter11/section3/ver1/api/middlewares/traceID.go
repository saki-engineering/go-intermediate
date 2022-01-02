package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func GetTraceID(ctx context.Context) int {
	id := ctx.Value("traceID")

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

func setTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}
