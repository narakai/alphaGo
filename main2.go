package main

import (
	"fmt"
)

//作用域

//有花括号"{ }"一般都存在作用域的划分； := 简式声明会屏蔽所有上层代码块中的变量（常量），
// 建议使用规则来规范，如对常量使用全部大写，而变量尽量小写； 在if等语句中存在隐式代码块，
// 需要注意； 闭包函数可以理解为一个代码块，并且他可使用包含它的函数内的变量；
var (
	Ga int = 99
)

const (
	v int = 199
)

func GetGa() func() int {

	if Ga := 55; Ga < 60 {
		fmt.Println("GetGa if 中：", Ga)
	}

	for Ga := 2; ; {
		fmt.Println("GetGa循环中：", Ga)
		break
	}

	fmt.Println("GetGa函数中：", Ga)

	return func() int {
		Ga += 1
		return Ga
	}
}

//func main() {
//	Ga := "string"
//	fmt.Println("main函数中：", Ga)
//
//	b := GetGa()
//	fmt.Println("main函数中：", b(), b(), b(), b())
//
//	v := 1
//	{
//		v := 2
//		fmt.Println(v)
//		{
//			v := 3
//			fmt.Println(v)
//		}
//	}
//	fmt.Println(v)
//}

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}
}
