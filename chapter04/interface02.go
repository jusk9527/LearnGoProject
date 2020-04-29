package main

//
//// 多态
//// 上面提到了，go 语言中interface 是实现多态的一种形式，所谓多态，就是一种事物的多种形态
//// 与python中类的多态是一致的
//
//// 同一个interface, 不同的类型实现，都可以进行调用，他们都按照统一接口进行操作
//// 在上面的实例中，我们增加一个Teacher 结构体，同样实现接口进行说明
//
//import "fmt"
//
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
//type Teacher struct {
//	Name   string
//	Salary int
//}
//
//func (p Student) Getname() string { // 实现Getname方法
//	fmt.Println(p.Name)
//
//	return p.Name
//
//}
//
//func (p Student) Running() { // 实现Runniing方法
//	fmt.Printf("%s running", p.Name)
//
//}
//
//func (p Teacher) Getname() string { // 实现Getname方法
//	fmt.Println(p.Name)
//	return p.Name
//}
//
//func (p Teacher) Running() {
//	fmt.Printf("\n%s running", p.Name)
//}
//
//func main() {
//	var skill Skills
//	var stu1 Student
//	var t1 Teacher
//
//	t1.Name = "wang"
//	stu1.Name = "wd"
//	stu1.Age = 22
//	skill = stu1
//	skill.Running()
//	skill = t1
//	t1.Running()
//
//}
//
////wd running
////wang running
