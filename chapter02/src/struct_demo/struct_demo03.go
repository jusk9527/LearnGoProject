package main

//
//// tag 可以为结构体的成员添加说明或者标签便于使用，这些说明可哦通过放射获取到
//// 在前面提到了，结构体中的成员首字母对外不可见，但是我们把成员定义为首自谋大写这样与外界进行数据交互会带来极大的不便
//// 此时tag带来了解决方法
//
//type Student struct {
//	Name  string "the name of student"
//	Age   int    "the age of student"
//	Class string "the class of student"
//}
