package main

//
//// 匿名成员
//// 结构体中，每个成员不一定都有名称，也允许字段没有名字，即匿名成员
//// 匿名成员的一个重要作用，可以用来实现oop中的继承
//// 同一种类型匿名成员值允许最多存在一个
//// 当匿名成员是结构体时，且两个结构体重都存在相同字段时，有限选择最近的字段
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//type Student struct {
//	score string
//	Age   int
//	Person
//}
//
//func main() {
//	var stu = new(Student)
//	stu.Age = 22
//	fmt.Println(stu.Person.Age, stu.Age)
//}
