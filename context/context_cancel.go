package main

import (
	"context"
	"fmt"
	"time"
)

func worker2(ctx context.Context)  {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("got the stop channel")
			return
		default:
			fmt.Println("still working")
			time.Sleep(1*time.Second)
		}
	}
}

func cancel_worker2(cancel context.CancelFunc){
	time.Sleep(5*time.Second)
	fmt.Println("stop the gorutine")
	cancel()
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)

	//go cancel_worker2(cancel) //cancel在goroutine里执行不会作用到worker里
	cancel_worker2(cancel)
	time.Sleep(5*time.Second)
}