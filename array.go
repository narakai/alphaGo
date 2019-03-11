package main

import (
	"fmt"
)

//数组
//var arrAge  = [5]int{18, 20, 15, 22, 16}
//var arrName = [5]string{3: "Chris", 4: "Ron"} //指定索引位置初始化
//// {"","","","Chris","Ron"}
//var arrCount = [4]int{500, 2: 100} //指定索引位置初始化 {500,0,100,0}
//var arrLazy = [...]int{5, 6, 7, 8, 22} //数组长度初始化时根据元素多少确定
//var arrPack = [...]int{10, 5: 100} //指定索引位置初始化，数组长度与此有关 {10,0,0,0,100}
//var arrRoom [20]int
//var arrBed = new([20]int)

//Go 语言中的数组是一种值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建：
//var arr1 = new([5]int)
//那么这种方式和 var arr2 [5]int 的区别是什么呢？arr1 的类型是 *[5]int，而 arr2的类型是 [5]int。

func main() {

	var arr1 = new([5]int)
	arr := arr1 //指针
	arr1[2] = 100
	fmt.Println(arr1[2], arr[2])

	var arr2 [5]int
	newarr := arr2 //拷贝
	arr2[2] = 100
	fmt.Println(arr2[2], newarr[2])

	//遍历
	//数组在声明时需要确定长度
	var arrAge = [5]int{18, 20, 15, 22, 16}
	//切片
	var arrAgeSlice = []int{18, 20, 15, 22, 16}
	fmt.Println(arrAge)
	fmt.Println(arrAgeSlice)
	for i, v := range arrAge {
		fmt.Printf("%d 的年龄： %d\n", i, v)
	}
}

//从上面代码结果可以看到，new([5]int)创建的是数组指针，arr其实和arr1指向同一地址，
// 故而修改arr1时arr同样也生效。而newarr是由arr2值传递（拷贝），故而修改任何一个都不会改变另一个的值。在写函数或方法时，如果参数是数组，需要注意参数长度不能过大。

//由于把一个大数组传递给函数会消耗很多内存（值传递），在实际中我们通常有两种方法可以避免这种现象：
//传递数组的指针
//使用切片 而通常使用切片是第一选择
