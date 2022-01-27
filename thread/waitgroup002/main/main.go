package main

import (
	"fmt"
	"sync"

	"Learning/golang-learning/thread/waitgroup002"
)

func main() {

	taskChan := make(chan waitgroup002.Task, 10)
	resultChan := make(chan int, 10)
	wait := &sync.WaitGroup{}

	go waitgroup002.InitTask(taskChan)
	go waitgroup002.DistributeTask(taskChan, wait, resultChan)
	result := waitgroup002.ProcessResult(resultChan)

	fmt.Println(result)

}
