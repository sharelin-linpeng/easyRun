package application

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/db"
)

type Application struct {
	Id          string `db:"id" json:"id"`
	AppName     string `db:"app_name" json:"appName"`
	AppWorkPath string `db:"app_workpath" json:"appWorkPath"`
	AppFile     string `db:"app_file" json:"appFile"`
}

type applicationService struct {
}

// 查询列表
func (applicationService) QueryList() []Application {
	sqlStr := "SELECT id, app_name, app_workpath,app_file FROM application"

	var appList []Application
	if err := db.DB.Select(&appList, sqlStr); err != nil {
		fmt.Printf("applicationService.QueryList, err:%v\n", err)
	}
	return appList

}

// 根据ID查询
func (applicationService) QueryById(id string) *Application {
	sqlStr := "SELECT id, app_name, app_workpath,app_file FROM application where id = ?"

	var app Application
	if err := db.DB.Get(&app, sqlStr, id); err != nil {
		fmt.Printf("applicationService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &app
}

// 根据ID查询
func (applicationService) Add(application Application) {
	sqlStr := "INSERT INTO application(id, app_name, app_workpath,app_file) VALUE( ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, application.Id, application.AppName, application.AppWorkPath, application.AppFile); err != nil {
		fmt.Printf("applicationService.Add(), err:%v\n", err)
	}
}

// 更新
func (applicationService) Update(application Application) {
	sqlStr := "UPDATE  application SET app_name = ? , app_workpath = ? , app_file = ? WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, application.AppName, application.AppWorkPath, application.AppFile, application.Id); err != nil {
		fmt.Printf("applicationService.Update(), err:%v\n", err)
	}
}

// 删除
func (applicationService) Delete(id string) {
	sqlStr := "DELETE  FROM application  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("applicationService.Delete(%s), err:%v\n", id, err)
	}
}

var ApplicationService = &applicationService{}
