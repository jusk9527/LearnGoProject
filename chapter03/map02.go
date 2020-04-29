package main

// map 可以嵌套，类似json 个格式，声明时候只需要将value 改为map,同样使用之前需要初始化每一层map, 实例

import "fmt"

func main() {
	a := make(map[string]map[string]string, 10) // 二层map嵌套，声明外层map并初始化
	a["key1"] = make(map[string]string)         // 二层map嵌套，声明外层map并初始化
	a["key1"]["key2"] = "a1"
	a["key1"]["key3"] = "b1"
	a["key1"]["key4"] = "c1"

	fmt.Println(a)

}
