// map操作
// 增删改查、求长度

//mymap["name"] = "wd"			// 创新或者更新
//delete(mymap, "name")			// 删除
//name := mymap["name"]			// 查询
//
//len(mymap)						// 求长度

// 测试key是否存在
package main

import "fmt"

func main() {
	a := make(map[string]string, 10)
	a["key1"] = "wd"
	val, ok := a["key1"] //ok为true时，代表有key

	if ok {
		fmt.Println(val)

	} else {
		fmt.Println("key1 is not exitst")
	}
}
