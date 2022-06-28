package task

import (
	"log"
	"time"
)

type taskExecutor struct {
	tasks chan ExecTask
	quit  chan int
}

func newTaskExecutor() *taskExecutor {
	executor := taskExecutor{}
	executor.tasks = make(chan ExecTask, 100)
	executor.quit = make(chan int)
	go func() {
		for {
			select {
			case task := <-executor.tasks:
				task.process()
			case <-time.After(3 * time.Second):
				log.Println("执行器等待任务")
			case <-executor.quit:
				return
			}
		}
	}()
	return &executor
}

func (taskExecutor *taskExecutor) ExecTask(task ExecTask) {
	taskExecutor.tasks <- task
}

func (taskExecutor *taskExecutor) Stop() {
	taskExecutor.quit <- 1
}

var TaskExecutor = newTaskExecutor()
