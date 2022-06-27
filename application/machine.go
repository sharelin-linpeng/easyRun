package application

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/db"
)

type Machine struct {
	// ID
	Id string `db:"id" json:"id"`
	// 机器ID
	Ip string `db:"ip" json:"ip"`
	// 登录名称
	LoginName string `db:"login_name" json:"loginName"`
	// 登录密码
	Password string `db:"password" json:"password"`
	// 机器环境
	Env string `db:"env" json:"env"`
}

type machineService struct {
}

// 查询列表
func (machineService) QueryList() []Machine {
	sqlStr := "SELECT id, ip, login_name,password,env FROM machine"

	var machines []Machine
	if err := db.DB.Select(&machines, sqlStr); err != nil {
		fmt.Printf("machineService.QueryList, err:%v\n", err)
	}
	return machines

}

// 根据ID查询
func (machineService) QueryById(id string) *Machine {
	sqlStr := "SELECT id, ip, login_name,password,env FROM machine where id = ?"

	var machine Machine
	if err := db.DB.Get(&machine, sqlStr, id); err != nil {
		fmt.Printf("machineService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &machine
}

// 根据ID查询
func (machineService) Add(machine Machine) {
	sqlStr := "INSERT INTO machine(id, ip, login_name,password,env) VALUE( ?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, machine.Id, machine.Ip, machine.LoginName, machine.Password, machine.Env); err != nil {
		fmt.Printf("machineService.Add(), err:%v\n", err)
	}
}

// 更新
func (machineService) Update(machine Machine) {
	sqlStr := "UPDATE  machine SET ip = ? , login_name = ? , password = ? ,env = ?  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, machine.Ip, machine.LoginName, machine.Password, machine.Env); err != nil {
		fmt.Printf("machineService.Update(), err:%v\n", err)
	}
}

// 删除
func (machineService) Delete(id string) {
	sqlStr := "DELETE  FROM machine  WHERE id = ?"
	if _, err := db.DB.Exec(sqlStr, id); err != nil {
		fmt.Printf("machineService.Delete(%s), err:%v\n", id, err)
	}
}

var MachineService = &machineService{}
