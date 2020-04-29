package main

// 多维数组

import "fmt"

func main() {
	var a [5][4][3]int // 三维数组
	a[1][1][0] = 1
	a[1][0][0] = 1
	a[1][1][1] = 1
	fmt.Println(a)
}
