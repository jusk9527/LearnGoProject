package main

// 应用场景实例，json 序列化操作
import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:":"name"`
	Age  int    `json:"age"`
}

func main() {
	var stu = Student{Name: "wd", Age: 22}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode failed err", err)
		return
	}
	fmt.Println(string(data))
}
