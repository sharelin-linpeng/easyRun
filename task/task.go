package task

import (
	"log"

	"github.com/sharelin-linpeng/easyRun/application"
)

type ExecTask interface {
	process() error
}

// 更新代码任务
type UpdateCodeTask struct {
	codeBranch *application.CodeBranch
}

func NewUpdateCodeTask(codeBranch *application.CodeBranch) *UpdateCodeTask {
	task := UpdateCodeTask{codeBranch}
	return &task
}

func (updateCodeTask *UpdateCodeTask) process() error {
	// TODO 更新分支代码
	log.Println("更新代码")
	return nil
}

// 构建代码任务
type BuildCodeTask struct {
	codeBranch *application.CodeBranch
}

func NewBuildCodeTask(codeBranch *application.CodeBranch) *BuildCodeTask {
	task := BuildCodeTask{codeBranch}
	return &task
}

func (updateCodeTask *BuildCodeTask) process() error {
	// TODO 构建分支代码
	log.Println("构建代码")
	return nil
}

// 发布应用任务
type PublishAppTask struct {
	publishInfo application.PublishBindingInfo
}

func NewPublishAppTask(publishInfo application.PublishBindingInfo) *PublishAppTask {
	task := PublishAppTask{publishInfo}
	return &task
}

func (publishAppTask *PublishAppTask) process() error {
	// TODO 发布应用
	log.Println("发布应用")
	return nil
}

// 复合任务
type CompositeTask struct {
	tasks []ExecTask
}

func NewCompositeTask() *CompositeTask {
	return &CompositeTask{}
}

func (compositeTask *CompositeTask) AddTask(task ExecTask) *CompositeTask {
	compositeTask.tasks = append(compositeTask.tasks, task)
	return compositeTask
}

func (compositeTask *CompositeTask) process() error {
	// TODO 发布应用
	for _, task := range compositeTask.tasks {
		if err := task.process(); err != nil {
			return err
		}
	}
	return nil
}
