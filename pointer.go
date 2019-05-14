package main

//在Go语言中，指针类型表示指向给定类型（称为指针的基础类型）的变量的所有指针的集合。
// 符号 * 可以放在一个类型前，如 *T，那么它将以类型T为基础，生成指针类型*T。未初始化指针的值为nil。例如：
//type Point3D struct{ x, y, z float64 }
//var pointer *Point3D
//var i *[4]int

//符号 * 可以放在一个指针前，如 (*pointer)，那么它将得到这个指针指向地址上所存储的值，这称为反向引用。
// 不过在Go语言中，(*pointer).x可以简写为pointer.x。
//
//对于任何一个变量 var， 表达式var == *(&var)都是正确的。

//内存管理中的内存区域一般包括堆内存（heap）和栈内存（stack）， 栈内存主要用来存储当前调用栈用到的简单类型数据，如string，bool，int，float 等。这些类型基本上较少占用内存，容易回收，因此可以直接复制，进行垃圾回收时也比较容易做针对性的优化。
// 而复杂的复合类型占用的内存往往相对较大，存储在堆内存中，垃圾回收频率相对较低，代价也较大，因此传引用或指针可以避免进行成本较高的复制操作，并且节省内存，提高程序运行效率。

//指针的使用方法：
//
//定义指针变量；
//
//为指针变量赋值；
//
//访问指针变量中指向地址的值；
//
//在指针类型前面加上*号来获取指针所指向的内容。

import "fmt"

func main() {
	var a, b int = 20, 30 // 声明实际变量
	var ptra *int         // 声明指针变量
	var ptrb *int = &b

	ptra = &a // 指针变量的存储地址

	fmt.Printf("a  变量的地址是: %x\n", &a)
	fmt.Printf("b  变量的地址是: %x\n", &b)

	// 指针变量的存储地址
	fmt.Printf("ptra  变量的存储地址: %x\n", ptra)
	fmt.Printf("ptrb  变量的存储地址: %x\n", ptrb)

	// 使用指针访问值
	fmt.Printf("*ptra  变量的值: %d\n", *ptra)
	fmt.Printf("*ptrb  变量的值: %d\n", *ptrb)
}

//new() 和 make() 都在堆上分配内存，但是它们的行为不同，适用于不同的类型。
//
//new() 用于值类型的内存分配，并且置为零值。 make() 只用于切片、字典以及通道这三种引用数据类型的内存分配和初始化。
//
//new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。 make(T) 返回类型T的值（不是* T）。

//通过SetFinalizer，如果GC自动运行或者手动运行GC，则都能及时销毁这些资源，释放占用的内存而避免内存泄漏。
//
//GC过程中重要的函数func SetFinalizer(obj interface{}, finalizer interface{})有两个参数，参数一：obj必须是指针类型。参数二：finalizer是一个函数，其参数类型是obj的类型，其没有返回值。
