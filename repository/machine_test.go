package repository

import (
	"fmt"
	"testing"

	"github.com/sharelin-linpeng/easyRun/common/config"
	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/common/jsonutil"
	"github.com/sharelin-linpeng/easyRun/entity"
)

func TestMachineInsert(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	app := entity.Machine{}
	app.Id = "123"
	app.Ip = "127.0.0.1"
	app.LoginName = "aaa"
	app.Password = "aaa111"
	app.Env = "test"
	MachineService.Add(app)
}

func TestMachineQueryId(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	app := MachineService.QueryById("123")
	fmt.Println(jsonutil.Obj2Json(app))
}

func TestMachineQueryList(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	applist := MachineService.QueryList()
	fmt.Println(jsonutil.Obj2Json(applist))
}

func TestMachineDelete(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	MachineService.Delete("123")
}
