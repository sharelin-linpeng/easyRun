package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type taskService struct {
}

// 自定义任务
func (taskService) QueryList() ([]entity.Task, error) {
	sqlStr := "SELECT id, parent_id, sequence,type FROM task where parent_id =" + entity.CONTAINER + "AND " + "type=" + entity.CONTAINER

	var dataList []entity.Task
	if err := db.DB.Select(&dataList, sqlStr); err != nil {
		fmt.Printf("taskService.QueryList, err:%v\n", err)
		return nil, err
	}
	return dataList, nil

}

// 根据ID查询
func (taskService) QueryById(id string) (*entity.Task, error) {
	sqlStr := "SELECT id, parent_id, sequence,type FROM task where id = ?"

	var data entity.Task
	if err := db.DB.Get(&data, sqlStr, id); err != nil {
		fmt.Printf("taskService.QueryById(%s), err:%v\n", id, err)
		return nil, err
	}
	return &data, nil
}

// 根据ID查询
func (taskService) Add(task entity.Task) error {
	sqlStr := "INSERT INTO task(id, parent_id, sequence,type) VALUE( ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, task.Id, task.ParentId, task.Sequence, task.Type); err != nil {
		fmt.Printf("taskService.Add(), err:%v\n", err)
		return err
	}
	return nil
}

// 更新
func (taskService) Update(task entity.Task) error {
	sqlStr := "UPDATE  task SET parent_id = ? , sequence = ? , type = ? WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, task.Id, task.ParentId, task.Sequence, task.Type, task.Id); err != nil {
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
