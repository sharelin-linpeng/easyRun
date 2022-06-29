package main

import (
	"os"

	"github.com/sharelin-linpeng/easyRun/common/config"
	"github.com/sharelin-linpeng/easyRun/common/db"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/controller"
)

func main() {
	config.LoadConfigApp(os.Args[1])
	db.InitMySQL()
	controller.InitController()
	server.InitHttpServer("8888")
}
