package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type publishInfoService struct {
}

// 查询列表
func (publishInfoService) QueryList() ([]entity.PublishBindingInfo, error) {
	sqlStr :=
		`SELECT 
		p.id, p.status, 
		a.id application_id, a.app_name, a.app_workpath, a.app_file,
		m.id machine_id, m.ip, m.login_name,m.password,m.env,
		c.id branch_id, c.branch_name, c.git_url, c.branch, c.dir, c.commond, c.repo_local
		FROM 
		publishInfo p,application a,machine m,codebranch c
		WHERE
			p.application_id = a.id and
			p.machine_id = m.id and
			p.branch_id = c.id
		`

	var appList []entity.PublishBindingInfo
	if err := db.DB.Select(&appList, sqlStr); err != nil {
		fmt.Printf("publishInfoService.QueryList, err:%v\n", err)
		return nil, err
	}
	return appList, nil

}

// 根据ID查询
func (publishInfoService) QueryById(id string) (*entity.PublishBindingInfo, error) {
	sqlStr :=
		`SELECT 
		p.id, p.status, 
		a.id application_id, a.app_name, a.app_workpath, a.app_file,
		m.id machine_id, m.ip, m.login_name,m.password,m.env,
		c.id branch_id, c.branch_name, c.git_url, c.branch, c.dir, c.commond, c.repo_local
		FROM 
		publishInfo p,application a,machine m,codebranch c
		WHERE
			p.application_id = ? and
			p.application_id = a.id and
			p.machine_id = m.id and
			p.branch_id = c.id
		`
	var app entity.PublishBindingInfo
	if err := db.DB.Get(&app, sqlStr, id); err != nil {
		fmt.Printf("publishInfoService.QueryById(%s), err:%v\n", id, err)
		return nil, err
	}
	return &app, nil
}

// 根据ID查询
func (publishInfoService) Add(data entity.PublishInfo) error {
	sqlStr := "INSERT INTO publishInfo (id, application_id, machine_id,branch_id,status) VALUE( ?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, data.Id, data.ApplicationId, data.MachineId, data.BranchId, data.Status); err != nil {
		fmt.Printf("publishInfoService.Add(), err:%v\n", err)
		return err
	}
	return nil
}

// 更新
func (publishInfoService) Update(data entity.PublishInfo) error {
	sqlStr := "UPDATE  publishInfo SET application_id = ? , machine_id = ? , branch_id = ?, status=? WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, data.ApplicationId, data.MachineId, data.BranchId, data.Status, data.Id); err != nil {
		fmt.Printf("publishInfoService.Update(), err:%v\n", err)
		return err
	}
	return nil
}

// 删除
func (publishInfoService) Delete(id string) error {
	sqlStr := "DELETE  FROM publishInfo  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("publishInfoService.Delete(%s), err:%v\n", id, err)
		return err
	}
	return nil
}

var PublishInfoService = &publishInfoService{}
