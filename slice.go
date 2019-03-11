package main

//切片（slice） 是对底层数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），所以切片是一个引用类型（和数组不一样）。
//绝对不要用指针指向 slice，切片本身已经是一个引用类型，所以它本身就是一个指针!
//var slice1 []type = arr1[start:end]
//var x = []int{2, 3, 5, 7, 11}
//v := make([]int, 10, 50) 这样分配一个有 50 个 int 值的数组，并且创建了一个长度为 10，容量为 50 的 切片 v，该切片指向数组的前 10 个元素。

import "fmt"

func main() {
	//声明切片的格式是： var identifier []type（不需要说明长度）。一个切片在未初始化之前默认为 nil，长度为 0。 var slice1 []type = arr1[start:end]
	sli := make([]int, 5, 10)
	fmt.Printf("切片sli长度和容量：%d, %d\n", len(sli), cap(sli))
	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	fmt.Println(newsli)

	//切片也可以用类似数组的方式初始化
	var x = []int{2, 3, 5, 7, 11}
	fmt.Printf("切片x长度和容量：%d, %d\n", len(x), cap(x))

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5] // a[low : high : max]  max-low的结果表示容量  high-low为长度
	fmt.Printf("切片t长度和容量：%d, %d\n", len(t), cap(t))

	//fmt.Println(t[2]) // panic ，索引不能超过切片的长度

	data := get()
	fmt.Println(len(data), cap(data), &data[0]) // 显示: 3 3 数组首字节地址

	//陈旧的切片 append()函数操作后，有没有生成新的slice需要看原有slice的容量是否足够
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 输出 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 输出 2 2 [2 3]
	for i := range s2 {
		s2[i] += 20
	}
	// s2的修改会影响到数组数据，s1输出新数据
	fmt.Println(s1) // 输出 [1 22 23]
	fmt.Println(s2) // 输出 [22 23]

	s2 = append(s2, 4) // append  s2容量为2，这个操作导致了slice s2扩容，会生成新的底层数组。

	for i := range s2 {
		s2[i] += 10
	}
	// s1 的数据现在是老数据，而s2扩容了，复制数据到了新数组，他们的底层数组已经不是同一个了。
	fmt.Println(len(s1), cap(s1), s1) // 输出3 3 [1 22 23]
	fmt.Println(len(s2), cap(s2), s2) // 输出3 4 [32 33 14]
}

//通过改变切片长度得到新切片的过程称之为切片重组 reslicing，做法如下：slice1 = slice1[0:end]，其中 end 是新的末尾索引（即长度）。
//
//当我们在一个slice基础上重新划分一个slice时，新的slice会继续引用原有slice的数组。如果你忘了这个行为的话，
//在你的应用分配大量临时的slice用于创建新的slice来引用原有数据的一小部分时，会导致难以预期的内存使用。
//为了避免这个陷阱，我们需要从临时的slice中使用内置函数copy()，拷贝数据（而不是重新划分slice）到新切片。

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 显示: 10000 10000 数组首字节地址
	res := make([]byte, 3)
	copy(res, raw[:3]) // 利用copy 函数复制，raw 可被GC释放
	return res
}
