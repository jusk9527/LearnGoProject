package main

//
//// 如果实现了结构体中的String 方法，在使用fmt打印时候会调用该方法，类似与Python中的__str__方法
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (self *Person) String() string {
//	return self.Name
//}
//
//func main() {
//	var person1 = new(Person)
//	person1.Name = "wd"
//	person1.Age = 22
//	fmt.Println(person1) //wd
//}
