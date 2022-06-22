package main

import (
	"fmt"

	"learning/golang-learning/thread/channel001"
)

const NUMBER = 10

func main() {

	workerNumbers := NUMBER
	taskChan := make(chan channel001.Task, 10)
	resultChan := make(chan int, 10)
	done := make(chan struct{}, 10)

	go channel001.InitTask(taskChan)
	channel001.DistributeTask(taskChan, done, workerNumbers)
	go channel001.CloseResult(done, resultChan, workerNumbers)

	result := channel001.ProcessResult(resultChan)
	fmt.Println(result)

}
