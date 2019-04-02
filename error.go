package main

import (
	"fmt"
)

func main() {
	fmt.Println("=========================")
	fmt.Println("return:", fun1())

	fmt.Println("=========================")
	fmt.Println("return:", fun2())
	fmt.Println("=========================")

	fmt.Println("return:", fun3())
	fmt.Println("=========================")

	fmt.Println("return:", fun4())
}

//上面函数签名中的 i 就是有名返回值，如果fun1()中定义了 defer 代码块，是可以改变返回值 i 的，函数返回语句return i 可以简写为 return 。
//在Go语言中，return 语句不是原子操作
func fun1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	// 规则二 defer执行顺序为先进后出

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
	}()

	// 规则三 defer可以读取有名返回值（函数指定了返回参数名）

	return 0 //这里实际结果为2。如果是return 100呢
}

func fun2() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
	}()
	return i
}

func fun3() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println(t)
	}()
	return t
}

func fun4() int {
	i := 8
	// 规则一 当defer被声明时，其参数就会被实时解析
	defer func(i int) {
		i = 99
		fmt.Println(i)
	}(i)
	i = 19
	return i
}
