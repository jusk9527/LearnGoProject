package main

//
//import "fmt"
//
//// 当有了结构的方法时候，我们可以自己定义其初始化方法，由于结构体时值类型，所以我们使用指针才能改变其存储的值
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (self *Person) init(name string, age int) {
//	self.Name = name
//	self.Age = age
//
//}
//
//func main() {
//	var person1 = new(Person)
//	person1.init("wd", 22)
//	// (&person1).init("wd",22)
//	fmt.Println(person1)
//}
