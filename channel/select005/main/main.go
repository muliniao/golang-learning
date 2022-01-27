package main

import "fmt"

func main() {

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

forEnd001:
	for {
		select {
		case <-ch1:
			fmt.Println("ch1 is calling")
		case <-ch2:
			fmt.Println("ch2 is calling")
		default:
			break forEnd001
		}
	}
	fmt.Println("forEnd01")

	for {
		select {
		case <-ch1:
			fmt.Println("ch1 is calling")
		case <-ch2:
			fmt.Println("ch2 is calling")
		default:
			goto forEnd002
		}
	}
forEnd002:

	fmt.Println("forEnd02")

}
