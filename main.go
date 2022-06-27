package main

import (
	"encoding/json"
	"fmt"

	"github.com/sharelin-linpeng/easyRun/application"
)

func main() {
	applist := application.ApplicationService.QueryList()
	for i := 0; i < len(applist); i++ {
		str, _ := json.Marshal(applist[i])
		fmt.Println(string(str))
	}
	app := application.ApplicationService.QueryById("4")
	str, _ := json.Marshal(app)
	fmt.Println(string(str))
}
