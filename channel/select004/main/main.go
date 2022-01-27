package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 5)
	out := make(chan int, 10)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6

	close(ch1)
	close(ch2)

ForEnd:
	for {
		select {
		case v1, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			out <- v1
		case v2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			out <- v2
		default:
			break ForEnd
		}
	}

	close(out)
	for v := range out {
		fmt.Println(v)
	}

}
