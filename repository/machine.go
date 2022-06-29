package repository

import (
	"fmt"

	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/entity"
)

type machineService struct {
}

// 查询列表
func (machineService) QueryList() []entity.Machine {
	sqlStr := "SELECT id, ip, login_name,password,env FROM machine"

	var machines []entity.Machine
	if err := db.DB.Select(&machines, sqlStr); err != nil {
		fmt.Printf("machineService.QueryList, err:%v\n", err)
	}
	return machines

}

// 根据ID查询
func (machineService) QueryById(id string) *entity.Machine {
	sqlStr := "SELECT id, ip, login_name,password,env FROM machine where id = ?"

	var machine entity.Machine
	if err := db.DB.Get(&machine, sqlStr, id); err != nil {
		fmt.Printf("machineService.QueryById(%s), err:%v\n", id, err)
		return nil
	}
	return &machine
}

// 根据ID查询
func (machineService) Add(machine entity.Machine) {
	sqlStr := "INSERT INTO machine(id, ip, login_name,password,env) VALUE( ?, ?, ?, ?, ?)"
	if _, err := db.DB.Exec(sqlStr, machine.Id, machine.Ip, machine.LoginName, machine.Password, machine.Env); err != nil {
		fmt.Printf("machineService.Add(), err:%v\n", err)
	}
}

// 更新
func (machineService) Update(machine entity.Machine) {
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
