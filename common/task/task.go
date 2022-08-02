package task

import (
	"log"
	"os"
	"os/exec"

	"github.com/sharelin-linpeng/easyRun/common/aes"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	http2 "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// 初始化代码任务
type InitCodeTask struct {
	codeBranch entity.CodeBranch
}

func NewInitCodeTask(codeBranch entity.CodeBranch) *InitCodeTask {
	task := InitCodeTask{codeBranch}
	return &task
}

func (initCodeTask *InitCodeTask) process() error {
	// TODO 初始化分支代码
	log.Println("初始化代码")

	codeBranch := initCodeTask.codeBranch
	user, _ := aes.DePwdCode(codeBranch.User)
	pass, _ := aes.DePwdCode(codeBranch.Auth)
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
	repository.CodeBranchService.Update(codeBranch)
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
	codeBranch := updateCodeTask.codeBranch

	// 代码未初始化，执行初始化
	if codeBranch.Status == int(entity.CODE_BRANCH_NOT_INIT) {
		initCodeTask := NewInitCodeTask(*codeBranch)
		initCodeTask.process()
	}

	r, err := git.PlainOpen(codeBranch.Dir)
	if err != nil {
		log.Println("拉取代码失败" + err.Error())
		codeBranch.Message = "拉取代码失败" + err.Error()
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		codeBranch.Message = "拉取代码失败" + err.Error()
		return err
	}
	user, _ := aes.DePwdCode(codeBranch.User)
	pass, _ := aes.DePwdCode(codeBranch.Auth)
	err = w.Pull(&git.PullOptions{
		ReferenceName: plumbing.ReferenceName("refs/heads/" + codeBranch.Branch),
		RemoteName:    "origin", Auth: &http2.BasicAuth{Username: string(user), Password: string(pass)}})
	if err != nil {
		log.Println("拉取代码失败" + err.Error())
		codeBranch.Message = err.Error()

		return err
	}
	ref, err := r.Head()
	if err != nil {
		log.Println("获取Head失败" + err.Error())
		codeBranch.Message = "获取Head失败" + err.Error()
		return err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Println("获取Hash失败" + err.Error())
		codeBranch.Message = "获取Hash失败" + err.Error()
		return err
	}
	codeBranch.Message = commit.Message
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
	codeBranch := updateCodeTask.codeBranch
	cmd := exec.Command("bash", "-c", codeBranch.Commond)
	cmd.Dir = codeBranch.Dir
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	err := cmd.Run()
	if err != nil {
		log.Printf("构建代码失败: %v\n", err)
		return err
	}
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
	publishInfo := publishAppTask.publishInfo
	log.Printf("开始发布应用[%s][%s] => [%s]:[%s]\n", publishInfo.AppName, publishInfo.AppFile, publishInfo.Ip, publishInfo.AppWorkPath)
	cmd := exec.Command("find", publishInfo.Dir, "-name", publishInfo.AppFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("发布应用失败: %v\n", err)
		return err
	}
	if len(output) > 0 {
		filePath := string(output)
		filePath = filePath[:len(filePath)-1]
		cmd = exec.Command("scp", "-r", filePath, publishInfo.LoginName+"@"+publishInfo.Ip+":"+publishInfo.AppWorkPath)
		log.Println("发布应用", cmd.Args)
		output, err = cmd.CombinedOutput()
		if err != nil {
			log.Printf("发布应用失败: %v\n", err)
			return err
		}
		cmd = exec.Command("ssh", publishInfo.LoginName+"@"+publishInfo.Ip, "cd", publishInfo.AppWorkPath+";./stop.sh")
		log.Println("发布应用", cmd.Args)
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		err := cmd.Run()
		if err != nil {
			log.Printf("发布应用失败: %v\n", err)
			return nil
		}

		cmd.Run()
		cmd = exec.Command("ssh", publishInfo.LoginName+"@"+publishInfo.Ip, "cd", publishInfo.AppWorkPath+";./start.sh")
		log.Println("发布应用", cmd.Args)
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		err = cmd.Run()
		if err != nil {
			log.Printf("发布应用失败: %v\n", err)
			return nil
		}
		cmd.Run()
	}
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
