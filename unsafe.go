package main

import (
	"fmt"
	"unsafe"
)

//ArbitraryType 是以int为基础定义的一个新类型，但是Go 语言unsafe包中，对ArbitraryType赋予了特殊的意义，通常，我们把interface{}看作是任意类型，那么ArbitraryType这个类型，在Go 语言系统中，比interface{}还要随意。
//
//Pointer 是ArbitraryType指针类型为基础的新类型，在Go 语言系统中，可以把Pointer类型，理解成任何指针的亲爹

//func Alignof(x ArbitraryType) uintptr
//func Offsetof(x ArbitraryType) uintptr
//func Sizeof(x ArbitraryType) uintptr

//Alignof返回变量对齐字节数量
//Offsetof返回变量指定属性的偏移量，所以如果变量是一个struct类型，不能直接将这个struct类型的变量当作参数，只能将这个struct类型变量的属性当作参数。
//Sizeof 返回变量在内存中占用的字节数，切记，如果是slice，则不会返回这个slice在内存中的实际占用长度。

//unsafe中，通过ArbitraryType 、Pointer 这两个类型，可以将其他类型都转换过来，然后通过这三个函数，分别能取长度，偏移量，对齐字节数，就可以在内存地址映射中，来回游走。

//只有将Pointer类型先转换成uintptr类型，做完地址加减法运算后，再转换成Pointer类型，通过*操作达到取值、修改值的目的。

//uintptr和unsafe.Pointer的区别就是：
//
//unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
//
//而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
//
//unsafe.Pointer 可以和 普通指针 进行相互转换；
//
//unsafe.Pointer 可以和 uintptr 进行相互转换。

type V struct {
	i int32
	j int64
}

func (v V) GetI() {
	fmt.Printf("i=%d\n", v.i)
}
func (v V) GetJ() {
	fmt.Printf("j=%d\n", v.j)
}

func main() {
	// 定义指针类型变量
	var v *V = &V{199, 299}

	// 取得v的指针并转为*int32的值，对应结构体的i。
	var i *int32 = (*int32)(unsafe.Pointer(v))

	fmt.Println("指针地址：", i)
	fmt.Println("指针uintptr值:", uintptr(unsafe.Pointer(i)))
	*i = int32(98)

	// 根据v的基准地址加上偏移量进行指针运算，运算后的值为j的地址，使用unsafe.Pointer转为指针
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0)))))

	*j = int64(763)

	v.GetI()
	v.GetJ()
}