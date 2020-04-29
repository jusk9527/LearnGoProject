package main

import "fmt"

// 数组遍历

func test(a *[3]int) { //传递地址
	(*a)[0] = 100
	return

}

func main() {
	var arr [3]int
	test(&arr) // 使用&取地址
	fmt.Println(arr[0])
}
