package main

// reflect.Type 中方法
// 同样方法

// 通用方法

//func (t *rtype) String() string // 获取 t 类型的字符串描述，不要通过 String 来判断两种类型是否一致。
//
//func (t *rtype) Name() string // 获取 t 类型在其包中定义的名称，未命名类型则返回空字符串。
//
//func (t *rtype) PkgPath() string // 获取 t 类型所在包的名称，未命名类型则返回空字符串。
//
//func (t *rtype) Kind() reflect.Kind // 获取 t 类型的类别。
//
//func (t *rtype) Size() uintptr // 获取 t 类型的值在分配内存时的大小，功能和 unsafe.SizeOf 一样。
//
//func (t *rtype) Align() int  // 获取 t 类型的值在分配内存时的字节对齐值。
//
//func (t *rtype) FieldAlign() int  // 获取 t 类型的值作为结构体字段时的字节对齐值。
//
//func (t *rtype) NumMethod() int  // 获取 t 类型的方法数量。
//
//func (t *rtype) Method() reflect.Method  // 根据索引获取 t 类型的方法，如果方法不存在，则 panic。
//// 如果 t 是一个实际的类型，则返回值的 Type 和 Func 字段会列出接收者。
//// 如果 t 只是一个接口，则返回值的 Type 不列出接收者，Func 为空值。
//
//func (t *rtype) MethodByName(string) (reflect.Method, bool) // 根据名称获取 t 类型的方法。
//
//func (t *rtype) Implements(u reflect.Type) bool // 判断 t 类型是否实现了 u 接口。
//
//func (t *rtype) ConvertibleTo(u reflect.Type) bool // 判断 t 类型的值可否转换为 u 类型。
//
//func (t *rtype) AssignableTo(u reflect.Type) bool // 判断 t 类型的值可否赋值给 u 类型。
//
//func (t *rtype) Comparable() bool // 判断 t 类型的值可否进行比较操作
//
//
//####注意对于：数组、切片、映射、通道、指针、接口
//func (t *rtype) Elem() reflect.Type // 获取元素类型、获取指针所指对象类型，获取接口的动态类型

import (
	"fmt"
	"reflect"
)

type Skills interface {
	reading()
	running()
}

type Student struct {
	Name string
	Age  int
}

func (self Student) runing() {
	fmt.Printf("%s is running\n", self.Name)
}

func (self Student) reading() {
	fmt.Printf("%s is reading\n", self.Name)

}

func main() {
	stu1 := Student{Name: "wd", Age: 22}
	inf := new(Skills)
	stu_type := reflect.TypeOf(stu1)
	inf_type := reflect.TypeOf(inf).Elem() // 特别说明，引用类型需要用Elem()获取指针所指的对象类型
	fmt.Println(stu_type.String())         // main.Student
	fmt.Println(stu_type.Name())           // Student
	fmt.Println(stu_type.PkgPath())        //main
	fmt.Println(stu_type.Kind())           // struct
	fmt.Println(stu_type.Size())           //24
	fmt.Println(inf_type.NumMethod())      // 2
	fmt.Println(inf_type.Method(0), inf_type.Method(0).Name)
	fmt.Println(inf_type.MethodByName("reading"))
}
