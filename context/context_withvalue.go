package main

import (
	"context"
	"fmt"
)

func main()  {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}

	k := favContextKey("language")
	//传入数据的类型和值k
	ctx := context.WithValue(context.Background(), k, "go")
	ctx2 := context.WithValue(context.Background(), k, "python")

	f(ctx, k)
	f(ctx2, k)
}
