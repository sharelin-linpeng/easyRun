package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type taskService struct {
}

// 自定义任务
func (taskService) QueryList() ([]entity.TaskVo, error) {
	sqlStr := `
	SELECT 
		id,
		name,
		parent_id, 
		sequence,
		type,
		code_branch_id,
		build_path,
		app_id,
		machine_id
	FROM task WHERE parent_id = 0`

	var dataList []entity.Task
	if err := db.DB.Select(&dataList, sqlStr); err != nil {
		fmt.Printf("taskService.QueryList, err:%v\n", err)
		return nil, err
	}

	taskVo := make([]entity.TaskVo, 0)
	for _, data := range dataList {
		task := entity.TaskVo{}
		task.Id = data.Id
		task.Name = data.Name
		childTaskList := make([]interface{}, 0)
		// 获取子任务
		if childTasks, err := TaskService.QueryById(data.Id); err != nil && childTasks != nil {
			for _, childTask := range childTasks {
				if childTask.Type == entity.UPDATE_CODE {

					// 查询分支信息
					if codeBranch, err := CodeBranchService.QueryById(childTask.CodeBranchId); err != nil {
						updateTask := entity.NewUpdateCodeTaskVo()
						updateTask.BranchName = codeBranch.BranchName
						updateTask.BranchId = codeBranch.Id
						childTaskList = append(childTaskList, updateTask)
					}

				} else if childTask.Type == entity.BUILD_CODE {

					// 查询分支信息
					if codeBranch, err := CodeBranchService.QueryById(childTask.CodeBranchId); err != nil {
						buildTask := entity.NewBuildCodeTaskVo()
						buildTask.BranchId = codeBranch.Id
						buildTask.BranchName = codeBranch.BranchName

						if childTask.BuildPath == "" {
							buildTask.BuildPath = codeBranch.Dir
						} else {
							buildTask.BuildPath = childTask.BuildPath
						}
						childTaskList = append(childTaskList, buildTask)
					}

				} else if childTask.Type == entity.PUBLISH_APP {
					publishTask := entity.NewPublishAppTaskVo()

					childTaskList = append(childTaskList, publishTask)
				}
			}
		}
		taskVo = append(taskVo, task)
	}

	return taskVo, nil

}

// 根据ID查询
func (taskService) QueryById(id int64) ([]entity.Task, error) {
	sqlStr := `
	SELECT 
		id,
		name,
		parent_id, 
		sequence,
		type,
		code_branch_id,
		build_path,
		app_id,
		machine_id
	FROM task WHERE parent_id = ? order by sequence`

	var data []entity.Task

	if err := db.DB.Select(data, sqlStr, id); err != nil {
		fmt.Printf("taskService.QueryById(%s), err:%v\n", id, err)
		return nil, err
	}
	return data, nil
}

// 根据ID查询
func (taskService) Add(task entity.Task) error {
	sqlStr := `
	INSERT INTO task(
		id,
		name,
		parent_id, 
		sequence,
		type,
		code_branch_id,
		build_path,
		app_id,
		machine_id) VALUE( ?,?, ?, ?, ?, ?, ?, ?)`
	if _, err := db.DB.Exec(sqlStr, task.Id, task.Name, task.ParentId, task.Sequence, task.Type, task.CodeBranchId, task.BuildPath, task.AppId, task.MachineId); err != nil {
		fmt.Printf("taskService.Add(), err:%v\n", err)
		return err
	}
	return nil
}

// 更新
func (taskService) Update(task entity.Task) error {
	sqlStr := `UPDATE  task SET 
	name = ?,
	parent_id = ? , 
	sequence = ? , 
	type = ? ,
	code_branch_id = ?,
	build_path = ?,
	app_id = ?,
	machine_id = ?
	WHERE id = ?`
	if _, err := db.DB.Exec(sqlStr, task.Name, task.ParentId, task.Sequence, task.Type, task.CodeBranchId, task.BuildPath, task.AppId, task.MachineId, task.Id); err != nil {
		fmt.Printf("taskService.Update(), err:%v\n", err)
		return err
	}
	return nil
}

// 删除
func (taskService) Delete(id string) error {
	sqlStr := "DELETE  FROM task  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("taskService.Delete(%s), err:%v\n", id, err)
		return err
	}
	return nil
}

var TaskService = &taskService{}
