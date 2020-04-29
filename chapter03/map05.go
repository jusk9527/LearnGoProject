package main

// 切片嵌套map

import "fmt"

func main() {
	a := map[string]string{"NAME": "WD", "AGE": "22"} // 初始化
	b := make([]map[string]string, 3, 3)              // 初始化切片

	b[0] = a
	fmt.Println(b)
	fmt.Println(a)
}
