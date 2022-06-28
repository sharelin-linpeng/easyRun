package task

import (
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	task := NewUpdateCodeTask(nil)
	TaskExecutor.ExecTask(task)
	time.Sleep(10 * time.Second)
}

func TestCompositeTask(t *testing.T) {
	compositeTask := NewCompositeTask()
	compositeTask.AddTask(NewUpdateCodeTask(nil))
	compositeTask.AddTask(NewBuildCodeTask(nil))
	TaskExecutor.ExecTask(compositeTask)
	time.Sleep(10 * time.Second)
}
