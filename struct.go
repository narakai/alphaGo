package main

import (
	"fmt"
	"reflect"
)

//结构体是值类型，因此可以通过 new 函数来创建。

//type identifier struct {
//	field1 type1
//	field2 type2
//	...
//}

//对于匿名字段，必须将匿名字段指定为类型名称T或指向非接口类型名称* T的指针，并且T本身可能不是指针类型
struct {
T1        // 字段名 T1
*T2       // 字段名 T2
P.T3      // 字段名 T3
*P.T4     // f字段名T4
x, y int    // 字段名 x 和 y
}

//使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
type S struct { a int; b float64 }
//new(S)
//new(S)为S类型的变量分配内存，并初始化（a = 0，b = 0.0），返回包含该位置地址的类型* S的值。

//我们一般的惯用方法是：t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。
//也可以这样写：var t T ，也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型T。
//在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）。

//使用点号符“.”可以获取结构体字段的值structname.fieldname。无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的表示法来引用结构体的字段
type myStruct struct { i int }
var v1 myStruct    // v是结构体类型变量
var p1 *myStruct   // p是指向一个结构体类型变量的指针
v1.i
p1.i

//结构体变量有下面几种初始化方式，前面一种按照字段顺序，后面两种则对应字段名来初始化赋值：
type Interval struct {
	start  int
	end   int
}

intr := Interval{0, 3}            (A)
intr := Interval{end:5, start:1}    (B)
intr := Interval{end:5}           (C)

//这里 Point3D{}以及 Line{origin, Point3D{y: -4, z: 12.3}}都是结构体字面量。
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }

origin := Point3D{}                      //  Point3D 是零值
line := Line{origin, Point3D{y: -4, z: 12.3}}  //   line.q.x 是零值

//表达式 new(Type) 和 &Type{} 是等价的。&struct1{a, b, c} 是一种简写，底层仍然会调用 new ()

//结构体类型和字段的命名遵循可见性规则，一个导出的结构体类型中有些字段是导出的，也即首字母大写字段会导出；另一些不可见，也即首字母小写为未导出，对外不可见

//通过参考应用可见性规则，如果结构体名不能导出，可使用 new 函数使用工厂方法的方法达到同样的目的
type bitmap struct {
	Size int
	data []byte
}

func NewBitmap(size int) *bitmap {
	div, mod := size/8, size%8
	if mod > 0 {
		div++
	}
	return &bitmap{size, make([]byte, div)}
}

//结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）。
// 它是一个附属于字段的字符串，可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有 reflect 包能获取它
import (
"fmt"
"reflect"
)

type Student struct {
	name string "学生名字"          // 结构体标签
	Age  int    "学生年龄"          // 结构体标签
	Room int    `json:"Roomid"` // 结构体标签
}

func main() {
	st := Student{"Titan", 14, 102}
	fmt.Println(reflect.TypeOf(st).Field(0).Tag)
	fmt.Println(reflect.TypeOf(st).Field(1).Tag)
	fmt.Println(reflect.TypeOf(st).Field(2).Tag)
}

type Human struct {
	name string
}

type Person1 struct {           // 内嵌
	Human
}

type Person2 struct {           // 内嵌， 这种内嵌与上面内嵌有差异
	*Human
}

type Person3 struct{             // 聚合
	human Human
}

//嵌入在结构体中广泛使用，在Go语言中如果只考虑结构体和接口的嵌入组合方式，一共有下面3种：
//1.在接口中嵌入接口:
type Writer interface{
	Write()
}

type Reader interface{
	Read()
}

type Teacher interface{
	Reader
	Writer
}

//2.在结构体中内嵌接口:
//初始化的时候，内嵌接口要用一个实现此接口的结构体赋值；或者定义一个新结构体，可以把新结构体作为receiver，
// 实现接口的方法就实现了接口（先记住这句话，后面在讲述方法时会解释），这个新结构体可作为初始化时实现了内嵌接口的结构体来赋值。
type Writer interface {
	Write()
}

type Author struct {
	name string
	Writer
}

func (a Author) Write() {
	fmt.Println(a.name, "  Write.")
}


// 定义新结构体，重点是实现接口方法Write()
type Other struct {
	i int
}

// 新结构体Other实现接口方法Write()，也就可以初始化时赋值给Writer 接口
func (o Other) Write() {
	fmt.Println(" Other Write.")
}

func main() {

	//  方法一：Other{99}作为Writer 接口赋值
	Ao := Author{"Other", Other{99}}
	Ao.Write()

	// 方法二：简易做法，对接口使用零值，可以完成初始化
	Au := Author{name: "Hawking"}
	Au.Write()
}

//程序输出：
//Other   Write.
//Hawking   Write.

//3.在结构体中嵌入结构体
//在结构体嵌入结构体很好理解，但不能嵌入自身值类型，可以嵌入自身的指针类型即递归嵌套。
//在初始化时，内嵌结构体也进行赋值；外层结构自动获得内嵌结构体所有定义的字段和实现的方法。

type Human struct {
	name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	string        // 匿名字段
}

type Student struct {
	Human     // 匿名字段
	Room  int // 教室
	int       // 匿名字段
}

func main() {
	//使用new方式
	stu := new(Student)
	stu.Room = 102
	stu.Human.name = "Titan"
	stu.Gender = "男"
	stu.Human.Age = 14
	stu.Human.string = "Student"

	fmt.Println("stu is:", stu)
	fmt.Printf("Student.Room is: %d\n", stu.Room)
	fmt.Printf("Student.int is: %d\n", stu.int) // 初始化时已自动给予零值：0
	fmt.Printf("Student.Human.name is: %s\n", stu.name) //  (*stu).name
	fmt.Printf("Student.Human.Gender is: %s\n", stu.Gender)
	fmt.Printf("Student.Human.Age is: %d\n", stu.Age)
	fmt.Printf("Student.Human.string is: %s\n", stu.string)

	// 使用结构体字面量赋值
	stud := Student{Room: 102, Human: Human{"Hawking", "男", 14, "Monitor"}}

	fmt.Println("stud is:", stud)
	fmt.Printf("Student.Room is: %d\n", stud.Room)
	fmt.Printf("Student.int is: %d\n", stud.int) // 初始化时已自动给予零值：0
	fmt.Printf("Student.Human.name is: %s\n", stud.Human.name)
	fmt.Printf("Student.Human.Gender is: %s\n", stud.Human.Gender)
	fmt.Printf("Student.Human.Age is: %d\n", stud.Human.Age)
	fmt.Printf("Student.Human.string is: %s\n", stud.Human.string)
}

//内嵌结构体的字段，可以逐层选择来使用，如stu.Human.name。如果外层结构体中没有同名的name字段，也可以直接选择使用，如stu.name。
//
//通过对结构体使用new(T)，struct{filed:value}两种方式来声明初始化，分别可以得到*T指针变量，和T值变量。
//
//从上面程序输出结果中stu is: &{{Titan 男 14 Student} 102 0} 可以得知，stu 是指针变量。但是程序在调用此结构体变量的字段时并没有使用到指针，
// 这是因为这里的 stu.name 相当于(*stu).name，这是一个语法糖，一般都使用stu.name方式来调用，但要知道有这个语法糖存在。