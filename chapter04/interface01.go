package main

//
//import "fmt"
//
//// 接口的使用不仅仅针对结构体，自定义类型、变量等等都可以实现接口
//// 如果一个接口没有任何办法，我们称之为空接口，由于空接口咩有方法，所以任何类型都实现了空接口
//// 要实现一个接口，必须实现该接口里面的所有方法
//// 定义接口
//type Skills interface {
//	Running()
//	Getname() string
//}
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//// 实现接口
//
//func (p Student) Getname() string { //实现Getname 方法
//	fmt.Println(p.Name)
//	return p.Name
//}
//
//func (p Student) Running() { // 实现Running方法
//	fmt.Printf("%s running", p.Name)
//
//}
//
//func main() {
//	var skill Skills
//	var stu1 Student
//	stu1.Name = "wd"
//	stu1.Age = 22
//	skill = stu1
//	skill.Running() // 调用接口
//}
