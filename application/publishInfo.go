package application

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/db"
)

type PublishInfo struct {
	// ID
	Id string `db:"id" json:"id"`
	// 应用ID
	ApplicationId string `db:"application_id" json:"applicationId"`
	// 机器ID
	MachineId string `db:"machine_id" json:"machineId"`
	// 关联分支
	BranchId string `db:"branch_id" json:"branchId"`
	// 状态
	Status string `db:"status" json:"status"`
}

type PublishBindingInfo struct {
	// ID
	Id string `db:"id" json:"id"`

	ApplicationId string `db:"application_id" json:"applicationId"`
	AppName       string `db:"app_name" json:"appName"`
	AppWorkPath   string `db:"app_workpath" json:"appWorkPath"`
	AppFile       string `db:"app_file" json:"appFile"`

	// ID
	BranchId string `db:"branch_id" json:"branchId"`
	// 分支名称
	BranchName string `db:"branch_name" json:"branchName"`
	// git地址
	GitUrl string `db:"git_url" json:"gitUrl"`
	// 分支名称
	Branch string `db:"branch" json:"branch"`
	// 代码存储路径
	Dir string `db:"dir" json:"dir"`
	// 构建命令
	Commond string `db:"commond" json:"commond"`
	// 本地仓库路径
	RepoLocal string `db:"repo_local" json:"repoLocal"`

	// ID
	MachineId string `db:"machine_id" json:"machineId"`
	// 机器ID
	Ip string `db:"ip" json:"ip"`
	// 登录名称
	LoginName string `db:"login_name" json:"loginName"`
	// 登录密码
	Password string `db:"password" json:"password"`
	// 机器环境
	Env string `db:"env" json:"env"`
}

type publishInfoService struct {
}

// 查询列表
func (publishInfoService) QueryList() []PublishBindingInfo {
	sqlStr :=
		`SELECT 
		p.id, p.status, 
		a.id application_id, a.app_name, a.app_workpath, a.app_file,
		m.id machine_id, m.ip, m.login_name,m.password,m.env,
		c.id BranchId, c.branch_name, c.git_url, c.branch, c.dir, c.commond, c.repo_local
		FROM 
		publishInfo p,application a,machine m,code_branch c
		WHERE
			p.application_id = a.id and
			p.machine_id = m.id and
			p.branchId = c.id
		`

	var appList []PublishBindingInfo
	if err := db.DB.Select(&appList, sqlStr); err != nil {
		fmt.Printf("publishInfoService.QueryList, err:%v\n", err)
	}
	return appList

}

// 根据ID查询
func (publishInfoService) QueryById(id string) *PublishBindingInfo {
	sqlStr :=
		`SELECT 
		p.id, p.status, 
		a.id application_id, a.app_name, a.app_workpath, a.app_file,
		m.id machine_id, m.ip, m.login_name,m.password,m.env,
		c.id BranchId, c.branch_name, c.git_url, c.branch, c.dir, c.commond, c.repo_local
		FROM 
		publishInfo p,application a,machine m,code_branch c
		WHERE
			p.application_id = ? and
			p.application_id = a.id and
			p.machine_id = m.id and
			p.branchId = c.id
		`
	var app PublishBindingInfo
	if err := db.DB.Get(&app, sqlStr, id); err != nil {
		fmt.Printf("publishInfoService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &app
}

// 根据ID查询
func (publishInfoService) Add(data PublishInfo) {
	sqlStr := "INSERT INTO publishInfo (id, application_id, machine_id,branch_id,status) VALUE( ?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, data.Id, data.ApplicationId, data.MachineId, data.BranchId, data.Status); err != nil {
		fmt.Printf("publishInfoService.Add(), err:%v\n", err)
	}
}

// 更新
func (publishInfoService) Update(data PublishInfo) {
	sqlStr := "UPDATE  publishInfo SET application_id = ? , machine_id = ? , branch_id = ?, status=? WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, data.ApplicationId, data.MachineId, data.BranchId, data.Status, data.Id); err != nil {
		fmt.Printf("publishInfoService.Update(), err:%v\n", err)
	}
}

// 删除
func (publishInfoService) Delete(id string) {
	sqlStr := "DELETE  FROM publishInfo  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("publishInfoService.Delete(%s), err:%v\n", id, err)
	}
}

var PublishInfoService = &publishInfoService{}
