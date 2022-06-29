package task

import (
	"log"
	"os"

	"github.com/sharelin-linpeng/easyRun/common/aes"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	http2 "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// 初始化代码任务
type InitCodeTask struct {
	codeBranch *entity.CodeBranch
}

func NewInitCodeTask(codeBranch *entity.CodeBranch) *InitCodeTask {
	task := InitCodeTask{codeBranch}
	return &task
}

func (initCodeTask *InitCodeTask) process() error {
	// TODO 初始化分支代码
	log.Println("初始化代码")
	codeBranch := initCodeTask.codeBranch
	user, _ := aes.DePwdCode("user")
	pass, _ := aes.DePwdCode("pass")
	_, err := git.PlainClone(codeBranch.Dir, false, &git.CloneOptions{
		URL:           codeBranch.GitUrl,
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + codeBranch.Branch),
		Auth:          &http2.BasicAuth{Username: string(user), Password: string(pass)},
	})
	if err != nil {
		log.Println(err.Error())
		codeBranch.Message = err.Error()
		codeBranch.Status = int(entity.CODE_BRANCH_NOT_INIT)
	} else {
		codeBranch.Message = "初始化完成"
		codeBranch.Status = int(entity.CODE_BRANCH_OK)
	}
	repository.CodeBranchService.Update(*codeBranch)
	return nil
}

// 更新代码任务
type UpdateCodeTask struct {
	codeBranch *entity.CodeBranch
}

func NewUpdateCodeTask(codeBranch *entity.CodeBranch) *UpdateCodeTask {
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
	codeBranch *entity.CodeBranch
}

func NewBuildCodeTask(codeBranch *entity.CodeBranch) *BuildCodeTask {
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
	publishInfo entity.PublishBindingInfo
}

func NewPublishAppTask(publishInfo entity.PublishBindingInfo) *PublishAppTask {
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

	for _, task := range compositeTask.tasks {
		if err := task.process(); err != nil {
			return err
		}
	}
	return nil
}
