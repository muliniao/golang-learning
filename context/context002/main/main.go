package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctxa, cancel001 := context.WithCancel(context.Background())
	fmt.Println(cancel001)
	go work(ctxa, "ctxa")

	ctxb, cancel002 := context.WithCancel(ctxa)
	fmt.Println(cancel002)
	go work(ctxb, "ctxb")

	// WithValue()
	//ctxc := context.WithValue(ctxb, "ctxc-key", "ctxc-value")
	//fmt.Println("----------")
	//go work(ctxc, "ctxc")

	ctxc, cancel003 := context.WithCancel(ctxb)
	fmt.Println(cancel003)
	go work(ctxc, "ctxc")

	ctxd, cancel004 := context.WithCancel(ctxc)
	fmt.Println(cancel004)
	go work(ctxd, "ctxd")

	time.Sleep(3 * time.Second)

	//cancel001()
	cancel002()
	//cancel003()
	//cancel004()

	time.Sleep(20 * time.Second)

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
