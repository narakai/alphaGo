package main

import (
	"fmt"
	"time"
)

//val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值

type field struct {
	name string
}

//用其他语言的概念的话,就是类的方法,调用的语法也类似 这个例子中work函数是Worker类的方法.
// 然后调用这个方法的对象以指针形式以参数w传入函数内 其实就是func work(w *Worker,done chan *Worker)的语法糖
//this
func (p *field) print() {
	fmt.Println(p.name)
}

func (p *field) print2() {
	fmt.Println(p.name + "1")
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		//time.Sleep(1 * time.Second)
		v.print2()
		go v.print()
		//v.print()
	}
	time.Sleep(3 * time.Second)
	// goroutines （可能）显示: three, three, three
}

//func print2(in string) {
//	fmt.Println(in)
//}

//func main() {
//	data := []string{"one", "two", "three"}
//
//	for _, v := range data {
//		//匿名函数
//		go func(in string) {
//			fmt.Println(in)  //当前的迭代变量作为匿名goroutine的参数。
//		}(v)
//		//go print2(v)
//	}
//
//	time.Sleep(3 * time.Second)
//	// goroutines输出: one, two, three
//}
