package repository

import (
	"fmt"
	"testing"

	"github.com/sharelin-linpeng/easyRun/common/config"
	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/common/jsonutil"
	"github.com/sharelin-linpeng/easyRun/entity"
)

func TestApplicationInsert(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	app := entity.Application{}
	app.Id = "123"
	app.AppName = "小zz"
	app.AppWorkPath = "猪窝"
	app.AppFile = "胡萝卜"
	ApplicationService.Add(app)
}

func TestApplicationQueryId(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	app, _ := ApplicationService.QueryById("123")
	fmt.Println(jsonutil.Obj2Json(app))
}

func TestApplicationQueryList(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	applist, _ := ApplicationService.QueryList()
	fmt.Println(jsonutil.Obj2Json(applist))
}

func TestApplicationDelete(t *testing.T) {
	config.LoadConfigApp("../config.yaml")
	db.InitMySQL()
	ApplicationService.Delete("123")
}
