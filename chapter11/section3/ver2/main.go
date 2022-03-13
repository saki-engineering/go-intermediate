package main

import (
	"context"
	"fmt"

	"mymodule/fuga"
	"mymodule/hoge"
)

func main() {
	ctx := context.Background()

	ctx = hoge.SetTraceID(ctx, 1)
	ctx = fuga.SetTraceID(ctx, 2)

	hogeTraceID := hoge.GetTraceID(ctx) // hoge.SetTraceIDでセットしたトレースIDが見たい
	fugaTraceID := fuga.GetTraceID(ctx) // fuga.SetTraceIDでセットしたトレースIDが見たい

	fmt.Println(hogeTraceID, fugaTraceID)
}
