package main

// 遍历map

import "fmt"

func main() {
	a := map[string]string{"NAME": "WD", "AGE": "22"}
	for k := range a { // 使用key进行遍历
		fmt.Println(a[k])

	}

	for k, v := range a { // 使用key,value 进行遍历
		fmt.Println(k, v)
	}
}
