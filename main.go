package main

import (
	"os"

	"github.com/sharelin-linpeng/easyRun/common/config"
	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/controller"
)

func main() {
	config.LoadConfigApp(os.Args[1])
	db.InitMySQL()
	controller.InitHttpServer("8888")
}
