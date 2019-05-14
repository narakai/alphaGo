package main

import (
	"fmt"
	"reflect"
)

//func TypeOf(i interface{}) Type
//type Type interface
//
//func ValueOf(i interface{}) Value
//type Value struct

//例如，x 被定义为：var x float64 = 3.4，那么 reflect.TypeOf(x) 返回 float64，reflect.ValueOf(x) 返回 3.4。
// 实际上，反射是通过检查一个接口的值，变量首先被转换成空接口

//Type主要有： Kind() 将返回一个常量，表示具体类型的底层类型 Elem()方法返回指针、数组、切片、map、通道的基类型，这个方法要慎用，如果用在其他类型上面会出现panic
//
//Value主要有： Type() 将返回具体类型所对应的 reflect.Type（静态类型） Kind() 将返回一个常量，表示具体类型的底层类型
//
//反射可以在运行时检查类型和变量，例如它的大小、方法和 动态 的调用这些方法。这对于没有源代码的包尤其有用。

type Student1 struct {
	name string
}

//func main() {
//
//	var a int = 50
//	v := reflect.ValueOf(a) // 返回Value类型对象，值为50
//	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int
//	fmt.Println(v, t, v.Type(), t.Kind())
//
//	var b [5]int = [5]int{5, 6, 7, 8}
//	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(),reflect.TypeOf(b).Elem()) // [5]int array int
//
//	var Pupil Student1
//	p := reflect.ValueOf(Pupil) // 使用ValueOf()获取到结构体的Value对象
//
//	fmt.Println(p.Type()) // 输出:Student
//	fmt.Println(p.Kind()) // 输出:struct
//
//}

//在Go语言中，类型包括 static type和concrete type. 简单说 static type是你在编码是看见的类型(如int、string)，concrete type是实际的类型，runtime系统看见的类型。
//
//Type()返回的是静态类型，而kind()返回的是concrete type。上面代码中，在int，数组以及结构体三种类型情况中，可以看到kind()，type()返回值的差异。

//通过反射可以修改原对象
//虽然反射可以越过Go语言的导出规则的限制读取结构体中未导出的成员，但不能修改这些未导出的成员。因为一个struct中只有被导出的字段才是settable的。

//type Student struct {
//	name string
//	Age  int
//}

//func main() {
//
//	var a int = 50
//	v := reflect.ValueOf(a) // 返回Value类型对象，值为50
//	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int
//	fmt.Println(v, t, v.Type(), t.Kind(), reflect.ValueOf(&a).Elem())
//	seta := reflect.ValueOf(&a).Elem() // 这样才能让seta保存a的值
//	fmt.Println(seta, seta.CanSet())
//	seta.SetInt(1000)
//	fmt.Println(seta)
//
//	var b [5]int = [5]int{5, 6, 7, 8}
//	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())
//
//	var Pupil Student = Student{"joke", 18}
//	p := reflect.ValueOf(Pupil) // 使用ValueOf()获取到结构体的Value对象
//
//	fmt.Println(p.Type()) // 输出:Student
//	fmt.Println(p.Kind()) // 输出:struct
//
//	setStudent := reflect.ValueOf(&Pupil).Elem()
//	//setStudent.Field(0).SetString("Mike") // 未导出字段，不能修改，panic会发生
//	setStudent.Field(1).SetInt(19)
//	fmt.Println(setStudent)
//
//}

// 结构体
type ss struct {
	int
	string
	bool
	float64
}

func (s ss) Method1(i int) string  { return "结构体方法1" }
func (s *ss) Method2(i int) string { return "结构体方法2" }

var (
	structValue = ss{ // 结构体
		20,
		"结构体",
		false,
		64.0,
	}
)

// 复杂类型
var complexTypes = []interface{}{
	structValue, &structValue, // 结构体
	structValue.Method1, structValue.Method2, // 方法
}

func main() {
	// 测试复杂类型
	for i := 0; i < len(complexTypes); i++ {
		PrintInfo(complexTypes[i])
	}
}

func PrintInfo(i interface{}) {
	if i == nil {
		fmt.Println("--------------------")
		fmt.Printf("无效接口值：%v\n", i)
		fmt.Println("--------------------")
		return
	}
	v := reflect.ValueOf(i)
	PrintValue(v)
}

func PrintValue(v reflect.Value) {
	fmt.Println("--------------------")
	// ----- 通用方法 -----
	fmt.Println("String             :", v.String())  // 反射值的字符串形式
	fmt.Println("Type               :", v.Type())    // 反射值的类型
	fmt.Println("Kind               :", v.Kind())    // 反射值的类别
	fmt.Println("CanAddr            :", v.CanAddr()) // 是否可以获取地址
	fmt.Println("CanSet             :", v.CanSet())  // 是否可以修改
	if v.CanAddr() {
		fmt.Println("Addr               :", v.Addr())       // 获取地址
		fmt.Println("UnsafeAddr         :", v.UnsafeAddr()) // 获取自由地址
	}
	// 获取方法数量
	fmt.Println("NumMethod          :", v.NumMethod())
	if v.NumMethod() > 0 {
		// 遍历方法
		i := 0
		for ; i < v.NumMethod()-1; i++ {
			fmt.Printf("    ┣ %v\n", v.Method(i).String())
			//			if i >= 4 { // 只列举 5 个
			//				fmt.Println("    ┗ ...")
			//				break
			//			}
		}
		fmt.Printf("    ┗ %v\n", v.Method(i).String())
		// 通过名称获取方法
		fmt.Println("MethodByName       :", v.MethodByName("String").String())
	}

	switch v.Kind() {
	// 结构体：
	case reflect.Struct:
		fmt.Println("=== 结构体 ===")
		// 获取字段个数
		fmt.Println("NumField           :", v.NumField())
		if v.NumField() > 0 {
			var i int
			// 遍历结构体字段
			for i = 0; i < v.NumField()-1; i++ {
				field := v.Field(i) // 获取结构体字段
				fmt.Printf("    ├ %-8v %v\n", field.Type(), field.String())
			}
			field := v.Field(i) // 获取结构体字段
			fmt.Printf("    └ %-8v %v\n", field.Type(), field.String())
			// 通过名称查找字段
			if v := v.FieldByName("ptr"); v.IsValid() {
				fmt.Println("FieldByName(ptr)   :", v.Type().Name())
			}
			// 通过函数查找字段
			v := v.FieldByNameFunc(func(s string) bool { return len(s) > 3 })
			if v.IsValid() {
				fmt.Println("FieldByNameFunc    :", v.Type().Name())
			}
		}
	}
}