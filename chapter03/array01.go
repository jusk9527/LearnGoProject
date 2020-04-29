package main

import "fmt"

// 数组遍历
func main() {
	var arr = [...]int{0: 2, 1: 4, 2: 8}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
} //结果是 2，4，8
