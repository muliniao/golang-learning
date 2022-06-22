package waitgroup_002

import "sync"

type Task struct {
}

func (t *Task) do() {

}

func InitTask(taskChan chan<- Task) {

	taskChan <- Task{}
	taskChan <- Task{}

	close(taskChan)
}

func DistributeTask(taskChan <-chan Task, wait *sync.WaitGroup, resultChan chan int) {

	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(resultChan)

}

func ProcessTask(t Task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

func ProcessResult(resultChan chan int) string {
	// process result
	return "final result"
}
