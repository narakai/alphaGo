package main

import (
	"fmt"
	"io"
)

//Go 语言中的所有类型包括自定义类型都实现了interface{}接口，
//这意味着所有的类型如string、 int、 int64甚至是自定义的struct类型都拥有interface{}的接口，这一点interface{}和Java中的Object类比较相似。

type AA struct {
	Books int
}

type B interface {
	f()
}

func (a AA) f() {
	fmt.Println("A.f() ", a.Books)
}

type I int

func (i I) f() {
	fmt.Println("I.f() ", i)
}

func main() {
	var a AA = AA{Books: 9}
	a.f()

	var b B = AA{Books: 99} // 接口类型可接受结构体A的值，因为结构体A实现了接口
	b.f()

	var i I = 199 // I是int类型引申出来的新类型
	i.f()

	var b2 B = I(299) // 接口类型可接受新类型I的值，因为新类型I实现了接口
	b2.f()
}

type Buffer funcType

//一个接口可以包含一个或多个其他的接口，但是在接口内不能嵌入结构体，也不能嵌入接口自身，否则编译会出错
type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

//我们可以使用类型断言（value, ok := element.(T)）来测试在某个时刻接口变量 varI 是否包含类型 T 的值
//varI 必须是一个接口变量，否则编译器会报错：invalid type assertion: varI.(T) (non-interface type (type of I) on left) 。
//更安全的方式是使用以下形式来进行类型断言:
var varI I
varI = T("Tstring")
if v, ok := varI.(T); ok { // 类型断言
fmt.Println("varI类型断言结果为：", v) // varI已经转为T类型
varI.f()
}

//接口类型向普通类型转换有两种方式：Comma-ok断言和Type-switch测试。
//通过Type-switch做类型判断:
// Type-switch做类型判断:
var value interface{}

switch str := value.(type) {
case string:
fmt.Println("value类型断言结果为string:", str)

case Stringer:
fmt.Println("value类型断言结果为Stringer:", str)

default:
fmt.Println("value类型不在上述类型之中")
}
//可以用 Type-switch 进行运行时类型分析，但是在 type-switch 时不允许有 fallthrough 。
// Type-switch让我们在处理未知类型的数据时，比如解析 json 等编码的数据，会非常方便。

//测试一个值是否实现了某个接口 Comma-ok断言
var varI I
varI = T("Tstring")
if v, ok := varI.(T); ok { // 类型断言
fmt.Println("varI类型断言结果为：", v) // varI已经转为T类型
varI.f()
}

//类型可以通过继承多个接口来提供像多重继承一样的特性
type ReaderWriter struct {
	io.Reader
	io.Writer
}