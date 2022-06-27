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

func (applicationService) QueryList() []Application {
	sqlStr := "SELECT id, app_name, app_workpath,app_file FROM application"

	var appList []Application
	if err := db.DB.Select(&appList, sqlStr); err != nil {
		fmt.Printf("applicationService.QueryList, err:%v\n", err)
	}
	return appList

}

func (applicationService) QueryById(id string) *Application {
	sqlStr := "SELECT id, app_name, app_workpath,app_file FROM application where id = ?"

	var app Application
	if err := db.DB.Get(&app, sqlStr, id); err != nil {
		fmt.Printf("applicationService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &app

}

var ApplicationService = &applicationService{}
