package main

//
//// 结构体中的方法
//
//// go 语言中的方法时通在特定类型的变量上，因此自定义的类型都可以有方法，不仅仅是在结构体重
//
//// go 中的方法和传统的类的方法不太一样，方法和类并非组织在一起，传统的oop方法和类放在你一个文件里面，
//// 而go 语言只要在同一个包里就可，而分散在不同文件里，go 的理念就是数据和实现分离，引用官方的说法 “Methods are not mixed with the data definition (the structs): they are orthogonal to types; representation(data) and behavior (methods) are independent”
//// 方法的调用通过recv.methodName() ，其访问控制也是1通过大小写区分
//
//// 方法定义，其中recv 代表方法作用的结构体
//
//// func (recv type) methodNam(parameter_list)(return_value_list){...}
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) Getname() string { // p代表结构体本身的实例，类似python 中的self,这里p可以写成self
//	fmt.Println(p.Name)
//	return p.Name
//
//}
//
//func main() {
//	var person1 = new(Person)
//	person1.Age = 22
//	person1.Name = "wd"
//	person1.Getname() //wd
//}
