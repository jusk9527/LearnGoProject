package main

import "fmt"

// len 计算长度（长度是指已经被父之过值的最大下标 +1）
// cap：求容量 (容量是指切片可容纳的最多元素个数)
// copy 拷贝切片
// append : 向切片中最加元素

// 使用append 向切片最加元素，如果长度没超过定义的切片的长度，返回原来的切片地址，如果超过了长度
// 切片会扩容进行重新分配地址

func main() {
	var slice2 []int = make([]int, 2, 3)
	fmt.Println(len(slice2), cap(slice2))
}
