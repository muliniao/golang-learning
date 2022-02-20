package main

import (
	"context"
	"fmt"
	"time"
)

type OtherContext struct {
	context.Context
}

func main() {

	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")

	ctxb, _ := context.WithDeadline(ctxa, time.Now().Add(3*time.Second))
	go work(ctxb, "work2")

	oc := OtherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "andes, pass from main")

	go workWithValue(ctxc, "work3")

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("main stop")

}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}
