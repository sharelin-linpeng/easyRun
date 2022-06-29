package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sharelin-linpeng/easyRun/common/config"
)

var DB *sqlx.DB

func InitMySQL() {
	db, err := sqlx.Open("mysql", config.CONFIG.App.MySql)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	DB = db
}
