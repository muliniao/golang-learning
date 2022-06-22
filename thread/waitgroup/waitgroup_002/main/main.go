package main

import (
	"fmt"
	"learning/golang-learning/thread/waitgroup/waitgroup_002"
	"sync"
)

func main() {

	taskChan := make(chan waitgroup_002.Task, 10)
	resultChan := make(chan int, 10)
	wait := &sync.WaitGroup{}

	go waitgroup_002.InitTask(taskChan)
	go waitgroup_002.DistributeTask(taskChan, wait, resultChan)
	result := waitgroup_002.ProcessResult(resultChan)

	fmt.Println(result)

}
