package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type applicationService struct {
}

// 查询列表
func (applicationService) QueryList() ([]entity.Application, error) {
	sqlStr := "SELECT id, app_name, app_workpath,app_file FROM application"

	var appList []entity.Application
	if err := db.DB.Select(&appList, sqlStr); err != nil {
		fmt.Printf("applicationService.QueryList, err:%v\n", err)
		return nil, err
	}
	return appList, nil

}

// 根据ID查询
func (applicationService) QueryById(id string) (*entity.Application, error) {
	sqlStr := "SELECT id, app_name, app_workpath,app_file FROM application where id = ?"

	var app entity.Application
	if err := db.DB.Get(&app, sqlStr, id); err != nil {
		fmt.Printf("applicationService.QueryById(%s), err:%v\n", id, err)
		return nil, err
	}
	return &app, nil
}

// 根据ID查询
func (applicationService) Add(application entity.Application) error {
	sqlStr := "INSERT INTO application(app_name, app_workpath,app_file) VALUE( ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, application.AppName, application.AppWorkPath, application.AppFile); err != nil {
		fmt.Printf("applicationService.Add(), err:%v\n", err)
		return err
	}
	return nil
}

// 更新
func (applicationService) Update(application entity.Application) error {
	sqlStr := "UPDATE  application SET app_name = ? , app_workpath = ? , app_file = ? WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, application.AppName, application.AppWorkPath, application.AppFile, application.Id); err != nil {
		fmt.Printf("applicationService.Update(), err:%v\n", err)
		return err
	}
	return nil
}

// 删除
func (applicationService) Delete(id string) error {
	sqlStr := "DELETE  FROM application  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("applicationService.Delete(%s), err:%v\n", id, err)
		return err
	}
	return nil
}

var ApplicationService = &applicationService{}
