package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sharelin-linpeng/easyRun/config"
)

var DB *sqlx.DB = initMySQL()

func initMySQL() *sqlx.DB {
	db, err := sqlx.Open("mysql", config.CONFIG.App.MySql)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return db
}
