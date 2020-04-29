package main

// go 语言中的map 都是无序的，并且无内置排序方法，所以如果我们想要对map 进行排序，我们需要自己实现

// 方法
// 先获取所有的key,将key进行排序
// 按照排序好的key进行遍历

import (
	"fmt"
	"sort"
)

func main() {
	a := map[string]string{"1": "a", "2": "b", "3": "c", "4": "d"}
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sort.Strings(keys) // 排序切片key
	fmt.Println(keys, len(keys))
	for _, val := range keys { // 循环key取值
		fmt.Println(a[val])

	}
}
