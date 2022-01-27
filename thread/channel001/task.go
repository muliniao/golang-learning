package channel001

type Task struct {
}

func (t *Task) do() {

}

func InitTask(taskChan chan<- Task) {
	taskChan <- Task{}
	taskChan <- Task{}

	close(taskChan)
}

func DistributeTask(taskChan <-chan Task, done chan struct{}, workerNumbers int) {
	for i := 0; i < workerNumbers; i++ {
		go ProcessTask(taskChan, done)
	}
}

func ProcessTask(taskChan <-chan Task, done chan struct{}) {
	for t := range taskChan {
		t.do()
	}
	done <- struct{}{}
}

func ProcessResult(resultChan chan int) string {
	// process result
	return "final result"
}

func CloseResult(done chan struct{}, resultChan chan int, workerNumbers int) {
	for i := 0; i < workerNumbers; i++ {
		<-done
	}
	close(done)
	close(resultChan)
}
