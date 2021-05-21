package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("sleep 1 Second")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}

}
