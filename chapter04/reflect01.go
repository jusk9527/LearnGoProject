package main

// 反射是程序执行时检查其所拥有的结构，尤其是类型的一种能力，这是元编程的一种形式。它同时
// 也是照成混淆的重要来源
// 每一个语言的反射模型都不同（同一时候很多语言根本不支持反射，Python 通过hasattr方法实现
// go 语言中的反射通过refect包实现，reflect包实现了运行时反射，允许程序操作任意类型的对象
// 在介绍反射之前先说明下reflect包中的两个数据类Type 和Value。

// Type
// Type :type 类型用来表示一个go类型
// 不是所有go 类型的Type 值都能使用所有方法，请参见每个方法的获取是引用限制，在调用有分类限定的方法时。应先使用Kind方法获知类型的分类
// 调用该分类不支持的方法会导致运行时的panic

// 获取Type 对象的方法
// func TypeOf(i interface{}) Type

import (
	"fmt"
	"reflect"
)

func main() {
	str := "wd"
	res_type := reflect.TypeOf(str)
	fmt.Println(res_type)
}

// string
