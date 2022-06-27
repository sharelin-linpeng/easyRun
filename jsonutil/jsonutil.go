package jsonutil

import (
	"encoding/json"
	"fmt"
)

func Obj2Json(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("obj2Json acc err %v\n", err)
	}
	return string(str)
}

func Json2Obj(jsonStr string, obj *interface{}) {
	if err := json.Unmarshal([]byte(jsonStr), obj); err != nil {
		fmt.Printf("json2Obj acc err %v\n", err)
	}
}
