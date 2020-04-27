// 单行注释
/* 多行
   注释 */

// 导入包的子句在每个源文件的开头
// Main比较特殊，他用来声明可执行文件，而不是一个库

package chapter01

// Import 语句声明了当前文件引用的包

import (
	"fmt" // Go语言标准库中的包
	"os"
	//"io/ioutil" // 包含一些输入输出函数
	//m "main"    //数学标准库，在此文件中别名为m
	//"net/http"  //一个web服务器包
	//"os"        //系统底层函数，如文件读写
	//"strconv"   //字符串转换
)

func mian() {
	// 往标准输出打印一行
	// 用包名fmt 限制打印函数
	fmt.Println("天坑欢迎你")

	// 调动当前包的另一个函数
	beyondHello()

}

// 函数可以以在括号里加参数
// 如果没有参数的话，也需要一个空括号

func beyondHello() {
	var x int // 声明变量，变量必须在使用之前声明
	x = 3     // 变量赋值

	// 可以用：=来偷懒，他自动将变量类型、声明和赋值都搞定了
	y := 4
	sum, prod := learnMultiple(x, y)       // 返回多个变量的函数
	fmt.Println("sum:", sum, "prod", prod) //简单输出
	learnTypes()                           //少于y 分钟，学的更多！
}

// 快看快看我是跨行注释_(:3),Go语言的函数可有多个参数和* 和多个*返回值。这在函数中， 'x','y' 是参数， 'sum','prod' 是返回值的标识符
// （可以理解为名字） 且类型为int

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // 返回两个值
}

// 内置变量类型和关键词
func learnTypes() {
	// 短声明给你所想。
	str := "少说话多读书!" // String类型

	s2 := `这是一个
可以换行的字符串` // 同样是String类型

	// 非ascii字符。Go使用UTF-8编码。
	g := 'Σ' // rune类型，int32的别名，使用UTF-8编码

	f := 3.14195 // float64类型，IEEE-754 64位浮点数
	c := 3 + 4i  // complex128类型，内部使用两个float64表示

	// Var变量可以直接初始化。
	var u uint = 7 // unsigned 无符号变量，但是实现依赖int型变量的长度
	var pi float32 = 22. / 7

	// 字符转换
	n := byte('\n') // byte是uint8的别名

	// 数组（Array）类型的大小在编译时即确定
	var a4 [4]int           // 有4个int变量的数组，初始为0
	a3 := [...]int{3, 1, 5} // 有3个int变量的数组，同时进行了初始化

	// Array和slice各有所长，但是slice可以动态的增删，所以更多时候还是使用slice。
	s3 := []int{4, 5, 9}    // 回去看看 a3 ，是不是这里没有省略号？
	s4 := make([]int, 4)    // 分配4个int大小的内存并初始化为0
	var d2 [][]float64      // 这里只是声明，并未分配内存空间
	bs := []byte("a slice") // 进行类型转换

	// 切片（Slice）的大小是动态的，它的长度可以按需增长
	// 用内置函数 append() 向切片末尾添加元素
	// 要增添到的目标是 append 函数第一个参数，
	// 多数时候数组在原内存处顺次增长，如
	s := []int{1, 2, 3}    // 这是个长度3的slice
	s = append(s, 4, 5, 6) // 再加仨元素，长度变为6了
	fmt.Println(s)         // 更新后的数组是 [1 2 3 4 5 6]

	// 除了向append()提供一组原子元素（写死在代码里的）以外，我们
	// 还可以用如下方法传递一个slice常量或变量，并在后面加上省略号，
	// 用以表示我们将引用一个slice、解包其中的元素并将其添加到s数组末尾。
	s = append(s, []int{7, 8, 9}...) // 第二个参数是一个slice常量
	fmt.Println(s)                   // 更新后的数组是 [1 2 3 4 5 6 7 8 9]

	p, q := learnMemory() // 声明p,q为int型变量的指针
	fmt.Println(*p, *q)   // * 取值

	// Map是动态可增长关联数组，和其他语言中的hash或者字典相似。
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	// 在Go语言中未使用的变量在编译的时候会报错，而不是warning。
	// 下划线 _ 可以使你“使用”一个变量，但是丢弃它的值。
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs
	// 通常的用法是，在调用拥有多个返回值的函数时，
	// 用下划线抛弃其中的一个参数。下面的例子就是一个脏套路，
	// 调用os.Create并用下划线变量扔掉它的错误代码。
	// 因为我们觉得这个文件一定会成功创建。
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "这句代码还示范了如何写入文件呢")
	file.Close()

	// 输出变量
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl() // 回到流程控制
}

// 和其他变成语言不同的是，go 支持有名称的变量返回值
// 声明返回值带上一个名字允许我们在韩珊瑚类的不同位置
// 只用写return 一个词就能将函数类指定名称的变量返回
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return
}

// Go 全面支持垃圾回收。Go 有指针，但是不支持指针运算
// 你会因为空指针而犯错，但是不会应为增加指针而犯错
func learnMemory() (p, q *int) {
	// 返回int型变量指针p和q
	p = new(int) //内置函数new 分配内存
	// 自动将分配的int 赋值0, p不在是空的了。
	s := make([]int, 20) //给20个int 变量分配一块内存
	s[3] = 7             // 赋值
	r := -2              // 声明另一个局部变量
	return &s[3], &r     // & 取地址

}

func learnFlowControl() {
	// If 需要花括号，括号就免了
	if true {
		fmt.Println("这句话肯定被执行了")
	}

	// 用go fmt 命令可以帮您格式化代码，所以不用怕被人吐槽代码风格了，
	// 也不用容忍别人的代码风格
	if false {
		// pout

	} else {
		// gloat
	}

	// 如果太多嵌套的if 语句，推荐使用switch
	x := 1
	switch x {
	case 0:
	case 1:
		//隐式调用break 语句，匹配上一个即停止

	case 2:
		// 不会运行

		// 和if 一样，for 也不用括号
		for x := 0; x < 3; x++ { //++ 自增
			fmt.Println("遍历", x)

		}

		// x 我在这里还是 1. 为什么

		// for 是go 里唯一的循环关键字，不过他还有很多变种

		for { // 死循环
			break    // 骗你的
			continue // 不会运行的
		}

		// 用range 可以枚举 array、slice、string、map、channel 等不同类型
		// 对于channe, range返回一个值
		// array、slice、string、map 等其他类型返回一对儿

		for key, value := range map[string]int{"onn": 1, "two": 2, "three": 3} {
			// 打印map中的每一个键值对
			fmt.Printf("索引：%s, 值为: %d\n", key, value)

		}

		// 如果你只是想要只，那就用前面讲的下划线扔掉没用的

	}
}
