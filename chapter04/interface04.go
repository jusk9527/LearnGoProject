package main

// 类型转换
// 由于接口是一般类型，当我们使用接口时候可额能不知道他是那个类型实现的，基本类型我们有对应的方法进行类型转换
// 换，当然接口类型也有蕾西转换

// 当然我么也可以用这个方式来进行类型的判断

// 转换方式：

//var s int
//var x interface
//
//x = s
//y ,ok := x.(int)			// 将interface 转为int,ok可省略，当时省略以后转换失败会报错，
// true 转换成功,false转换失败，并采用默认值

import "fmt"

func main() {
	var x interface{}

	s := "WD"
	x = s
	y, ok := x.(int)
	z, ok1 := x.(string)

	fmt.Println(y, ok)
	fmt.Println(z, ok1)
}
