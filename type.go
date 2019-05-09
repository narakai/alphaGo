package main

import "fmt"

type A struct {
	Face int
}

//自定义类型不会拥有原类型附带的方法，而别名是拥有原类型附带的
type Aa A    // 自定义新类型Aa，没有基础类型A的方法
type Aaa = A // 类型别名

func (a A) f() {
	fmt.Println("hi ", a.Face)
}

func main() {
	var s A = A{Face: 9}
	s.f()

	//var sa Aa = Aa{ Face: 9 }
	//sa.f()

	var sa Aaa = Aaa{Face: 9}
	sa.f()
}

//自定义类型不会继承原有类型的方法，但接口方法或组合类型的内嵌元素则保留原有的方法。
//  Mutex 用两种方法，Lock and Unlock。
type Mutex struct { /* Mutex fields */
}

func (m *Mutex) Lock()   { /* Lock implementation */ }
func (m *Mutex) Unlock() { /* Unlock implementation */ }

// NewMutex和 Mutex 一样的数据结构，但是其方法是空的。
type NewMutex Mutex

// PtrMutex 的方法也是空的
type PtrMutex *Mutex

// *PrintableMutex 拥有Lock and Unlock 方法
type PrintableMutex struct {
	Mutex
}
