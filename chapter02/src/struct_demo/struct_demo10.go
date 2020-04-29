package main

//
//// 内存分布
//
//// go中的结构体内存布局和c结构体布局类似。每个成员的内存分布式连续的，在以下实例中通过哦反射进一步说明
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type Student struct {
//	Name  string
//	Age   int64
//	wight int64
//	high  int64
//	score int64
//}
//
//func main() {
//	var stu1 = new(Student)
//	fmt.Printf("%p\n", &stu1.Name)
//	fmt.Printf("%p\n", &stu1.Age)
//	fmt.Printf("%p\n", &stu1.wight)
//	fmt.Printf("%p\n", &stu1.high)
//	fmt.Printf("%p\n", &stu1.score)
//
//}
