package main

//map是一种元素对的无序集合，一组称为元素value，另一组为唯一键索引key。 未初始化map的值为nil。map 是引用类型
//var map1 map[keytype]valuetype
//key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。
//所以数组、函数、字典、切片和结构体不能作为 key (含有数组切片的结构体不能作为 key，
// 只包含内建类型的 struct 是可以作为 key 的），但是指针和接口类型可以。
//map 是引用类型的，内存用 make 方法来分配。map 的初始化: var map1 = make(map[keytype]valuetype)

//一般判断是否某个key存在，不使用值判断，而使用下面的方式：

//if _, ok := x["two"]; !ok {
//fmt.Println("no entry")
//}

//如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包）。

//在"range"语句中生成的数据的值是真实集合元素的拷贝，它们不是原有元素的引用。
// 这意味着更新这些值将不会修改原来的数据。同时也意味着使用这些值的地址将不会得到原有数据的指针。

import "fmt"

func main() {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10 // 通常数据项不会改变
	}

	fmt.Println("data:", data) // 程序输出: [1 2 3]

	for i, _ := range data {
		data[i] *= 10
	}

	fmt.Println("data:", data)
}

//程序输出：
//data: [1 2 3]

//如果你需要更新原有集合中的数据，使用索引操作符来获得数据。
