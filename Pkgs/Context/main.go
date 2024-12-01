package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//some examples of using context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()
	key := "key"
	value := "value"

	ctx = context.WithValue(ctx, key, value)
	value1 := ctx.Value(key)
	fmt.Println(value1)
	select {
	case <-ctx.Done():
		// do something
	default:
		// do something
	}

}
