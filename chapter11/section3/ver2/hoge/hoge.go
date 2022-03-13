package hoge

import "context"

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
