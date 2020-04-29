package main

//
//// 切片 是对数组的引用，数组的长度不便，但是切片长度是可以变化的
//// 切片的特征和数组一样，因此遍历，长度计算等和数组一样
//// 切片定义和数组的定义方式在于是否有长度定义，有长度定义为数组，没有长度定义为切片
//
//// 方式一，使用make创建
////slice1 := make([] type, len,capacity)		或者var slice1 [] type = make([]type, len,capacity)
//
//
//// 方式二，直接初始化
//slice2 :=[] int {1,2}			//长度为2，容量为2
//
//// 方式三，从数组中获取切片没有python 一样
//slice := arr[startIndex:endIndex]			// 将arr数组中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
//slice := arr[startIndex:]					// 缺省endIndex 时表示一直到arr数组的最后一个元素
//slice := arr[:endIndex]						// 缺省startIndex 时将表示从arr数组的第一个元素开始
//
