package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//https://github.com/ffhelicopter/Go42/blob/master/SUMMARY.md
func main() {
	const s = "Go语言"
	for i, r := range s {
		fmt.Printf("%#U  ： %d\n", r, i)
	}

	x := 1
	fmt.Println(x) // prints 1
	{
		fmt.Println(x) // prints 1
		x := 2
		fmt.Println(x) // prints 2
	}
	fmt.Println(x) // prints 1 (不是2)

	//_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，
	// 但有时你并不需要使用从一个函数得到的所有返回值。

	//由于Go语言有个强制规定，在函数内一定要使用声明的变量，但未使用的全局变量是没问题的。为了避免有未使用的变量，代码将编译失败，
	// 我们可以将该未使用的变量改为 _。

	//在Go语言中，如果引入的包未使用，也不能通过编译。有时我们需要引入的包，比如需要init()，或者调试代码时我们可能去掉了某些包的功能使用，
	// 你可以添加一个下划线标记符，_，来作为这个包的名字，从而避免编译失败。下滑线标记符用于引入，但不使用。
	var b int
	_, b = 5, 7
	fmt.Println(b)

	//并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到
	//val, err = Func1(var1)

	//在 Go 语言中，指针属于引用类型，其它的引用类型还包括 slices，maps和 channel。

	//Go中的数组是数值，因此当你向函数中传递数组时，函数会得到原始数组数据的一份复制。如果你打算更新数组的数据，可以考虑使用数组指针类型。
	y := [3]int{1, 2, 3}
	fmt.Println(y) // prints [1 2 3]

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) // prints &[7 2 3]
	}(&y)

	fmt.Println(y) // prints [7 2 3]

	//使用位左移与 iota 计数配合可优雅地实现存储单位的常量枚举：
	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	//简式变量 使用 := 定义的变量，如果新变量Ga与那个同名已定义变量 (这里就是那个全局变量Ga)不在一个作用域中时，
	// 那么Go 语言会新定义这个变量Ga，遮盖住全局变量Ga。刚开始很容易在此犯错而茫然，解决方法是局部变量尽量不同名。

	//根据 Go语言的规范 ，Go的标识符作用域是基于代码块（code block）的。代码块就是包裹在一对大括号内部的声明和语句，
	//并且是可嵌套的。在代码中直观可见的显式的(explicit)code block，
	//比如：函数的函数体、for循环的循环体等；还有隐式的(implicit)code block。

	//if simplestmt; expression {
	//	... ...
	//}
	//{ // 隐式code block
	//	simplestmt
	//	if expression { // 显式的code block
	//		... ...
	//	}
	//}

	//当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
	// 那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），
	// 这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，
	// 但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。

	//Go语言中init函数用于包(package)的初始化，该函数是Go语言的一个重要特性，有下面的特征：
	//
	//init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
	//每个包可以拥有多个init函数
	//包的每个源文件也可以拥有多个init函数
	//同一个包中多个init函数的执行顺序Go语言没有明确的定义(说明)
	//不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
	//init函数不能被其他函数调用，而是在main函数执行之前，自动被调用

	//双引号中的转义字符被替换，而反引号中原生字符串中的 \n 会被原样输出。

	//Go 语言中的string类型是一种值类型，存储的字符串是不可变的，
	//如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的。

	st := "其实就是rune"
	fmt.Println(len(st))                    // "16"
	fmt.Println(utf8.RuneCountInString(st)) // "8"

	//字符串拼接
	var b1 strings.Builder
	b1.WriteString("ABC")
	b1.WriteString("DEF")

	fmt.Print(b1.String())

	//标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
	//
	//strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
	//
	//bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效，稍后我们将展示。
	//
	//strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
	//
	//unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。
}
