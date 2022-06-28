package main

import (
	"os"

	"github.com/sharelin-linpeng/easyRun/config"
	"github.com/sharelin-linpeng/easyRun/controller"
	"github.com/sharelin-linpeng/easyRun/db"
)

func main() {
	config.LoadConfigApp(os.Args[1])
	db.InitMySQL()
	controller.InitHttpServer("8888")
}
