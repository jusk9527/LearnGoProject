package main

import "fmt"

// map 是go 语言内置的key - value 的数据结构，与python 的dict 类似，可称为字典或关联数组

// map 属于引用类型，声明是不会分配内存的，需要make 初始化分配内存

//只声明不初始化，直接使用会panic，需要使用make分配内存后方可使用
//var a map[keytype]valuetype
//var a map[string]string
//var a map[string]int
//var a map[int]string
//var a map[string]map[string]string

//声明并初始化
//var a map[string]string
//a = make(map[string]string, 8) //8代表容量

//a := make(map[string]string, 8)
//a := make(map[string]string)

//var a map[string]string = map[string]string{}
var a map[string]string = map[string]string{"name": "wd", "age": "22"}

func main() {
	fmt.Println("a的值", a)
}
