package main

//func 函数名(参数列表) (返回值列表) {
//	// 函数体
//	return
//}

import (
	"fmt"
	"time"
)

//函数也可以作为函数类型被使用。函数类型也就是函数签名，函数类型表示具有相同参数和结果类型的所有函数的集合。函数类型的未初始化变量的值为nil
type funcType func(time.Time) // 定义函数类型funcType

//func main() {
//	f := func(t time.Time) time.Time { return t } // 方式一：直接赋值给变量
//	fmt.Println(f(time.Now()))
//
//	var timer funcType = CurrentTime // 方式二：定义函数类型funcType变量timer
//	timer(time.Now())
//
//	funcType(CurrentTime)(time.Now()) // 先把CurrentTime函数转为funcType类型，然后传入参数调用
//	// 这种处理方式在Go 中比较常见
//
//}

func CurrentTime(start time.Time) {
	fmt.Println(start)
}

//Go 语言中函数默认使用按值传递来传递参数，也就是传递参数的副本。
// 函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量。

//如果我们希望函数可以直接修改参数的值，而不是对参数的副本进行操作，则需要将参数的地址传递给函数，
// 这就是按引用传递，比如 Function(&arg1)，此时传递给函数的是一个指针。如果传递给函数的是一个指针，我们可以通过这个指针来修改对应地址上的变量值。

//在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）等这样的引用类型都是默认使用引用传递。

//Go语言中函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调
//func main() {
//	callback(1, Add)
//}

func Add(a, b int) {
	fmt.Printf("%d 与 %d 相加的和是: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // 回调函数f
}

//匿名函数
//func main() {
//	fn := func() {
//		fmt.Println("hello")
//	}
//	fn()
//
//	fmt.Println("匿名函数加法求和：", func(x, y int) int { return x + y }(3, 4))
//
//	func() {
//		sum := 0
//		for i := 1; i <= 1e6; i++ {
//			sum += i
//		}
//		fmt.Println("匿名函数加法循环求和：", sum)
//	}()
//}

//匿名函数同样也被称之为闭包。
//
//闭包可被允许调用定义在其环境下的变量，可以访问它们所在的外部函数中声明的所有局部变量、参数和声明的其他内部函数。
// 闭包继承了函数所声明时的作用域，作用域内的变量都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁。
// 也可以理解为内层函数引用了外层函数中的变量或称为引用了自由变量。
//
//实质上看，闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。
// 闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。
// 由闭包的实质含义，我们可以推论：闭包获取捕获变量相当于引用传递，而非值传递；对于闭包函数捕获的常量和变量，无论闭包何时何处被调用，闭包都可以使用这些常量和变量，而不用关心它们表面上的作用域。

var G int = 7

//func main() {
//	// 影响全局变量G，代码块状态持续
//	y := func() int {
//		fmt.Printf("G: %d, G的地址:%p\n", G, &G)
//		G += 1
//		return G
//	}
//	fmt.Println(y(), &y)
//	fmt.Println(y(), &y)
//	fmt.Println(y(), &y) //y的地址
//
//	fmt.Println("-----------------")
//
//	// 影响全局变量G，注意z的匿名函数是直接执行，所以结果不变
//	z := func() int {
//		G += 1
//		return G
//	}()
//	fmt.Println(z, &z)
//	fmt.Println(z, &z)
//	fmt.Println(z, &z)
//
//	fmt.Println("-----------------")
//
//	// 影响外层（自由）变量i，代码块状态持续
//	var f = N()
//	fmt.Println(f(1), &f)
//	fmt.Println(f(1), &f)
//	fmt.Println(f(1), &f)
//
//	fmt.Println("-----------------")
//
//	var f1 = N()
//	fmt.Println(f1(1), &f1)
//
//}

func N() func(int) int {
	var i int
	return func(d int) int {
		fmt.Printf("i: %d, i的地址:%p\n", i, &i)
		i += d
		return i
	}
}

//G是闭包中被捕获的全局变量，因此，对于每一次引用，G的地址都是固定的，i是函数内部局部变量，地址也是固定的，他们都可以被闭包保持状态并修改。
// 还要注意，f和f1是不同的实例，它们的地址是不一样的。

//变参函数
func Greeting(who ...string) {
	for k, v := range who {

		fmt.Println(k, v)
	}
}

func main() {
	s := []string{"James", "Jasmine"}
	Greeting(s...) // 注意这里切片s... ，把切片打散传入，与s具有相同底层数组的值。
}
