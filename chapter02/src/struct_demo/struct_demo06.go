package main

//
//// 继承、多继承
//
//// 当结构体中的成员也是结构体时，该结构体就继承了这个结构体，继承了其所有的方法与属性，
//// 当然有多个结构体成员也就是多继承
//
//// 访问父结构中属性也使用".",但是当子结构体中存在和父结构中的字段相同的时候，只能使用："子结构体.父结构体.字段"访问福结构
//// 体中的属性，如上面实例的 stu.Person.Age
//
//// 继承结构体可以使用别名，访问的时候通过别名访问，如下面man1.job.Salary
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//type Teacher struct {
//	Salary  int
//	Classes string
//}
//
//type man struct {
//	sex    string
//	job    Teacher // 别名，继承Teacher
//	Person         // 继承Person
//
//}
//
//func main() {
//	var man1 = new(man)
//	man1.Age = 22
//	man1.Name = "wd"
//	man1.job.Salary = 8500
//	fmt.Println(man1, man1.job.Salary) //${{8500} {wd 22}} 8500
//}
