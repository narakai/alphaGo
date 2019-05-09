package main

import (
	"fmt"
	"reflect"
)

//Go 语言中方法和函数在形式上很像，它是作用在接收器（receiver）上的一个函数，接收器是某种类型的变量。
// 因此方法是一种特殊类型的函数，方法只是比函数多了一个接收器（receiver），当然在接口中定义的函数我们也称为方法（因为最终还是要通过绑定到类型来实现）。
//
//正是因为有了接收器，方法才可以作用于接收器的类型（变量）上，类似于面向对象中类的方法可以作用于类属性上。
//
//定义方法的一般格式如下：
//
//func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

//type A struct {
//	Face int
//}
//
//func (a A) f() {
//	fmt.Println("hi ", a.Face)
//}

//上面代码中，我们定义了结构体 A ，注意f()就是 A 的方法，(a A)表示接收器。a 是 A的实例，
// f()是它的方法名，方法调用遵循传统的 object.name 即选择器符号：a.f()。

//接收器类型除了不能是指针类型或接口类型外，可以是其他任何类型，不仅仅是结构体类型，也可以是函数类型，还可以是 int、bool、string 等等为基础的自定义类型。

type Human struct {
	name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	string        // 匿名字段
}

func (h Human) print() { // 值方法
	fmt.Println("Human:", h)
}

type MyInt int

func (m MyInt) print() { // 值方法
	fmt.Println("MyInt:", m)
}

func main() {
	//使用new方式
	hu := new(Human)
	hu.name = "Titan"
	hu.Gender = "男"
	hu.Age = 14
	hu.string = "Student"
	hu.print()

	// 指针变量
	mi := new(MyInt)
	mi.print()

	// 使用结构体字面量赋值
	hum := Human{"Hawking", "男", 14, "Monitor"}
	hum.print()

	// 值变量
	myi := MyInt(99)
	myi.print()
}

//在Go语言中，方法调用的方式如下：如有类型X的变量x，m()是其方法，则方法有效调用方式是x.m()，而x如果是指针变量，则 x.m() 是 (&x).m()的简写。
// 所以我们看到指针方法的调用往往也写成 x.m()，其实是一种语法糖。

//函数和方法的区别:
//方法相对于函数多了接收器，这是他们之间最大的区别
//函数是直接调用，而方法是作用在接收器上，方法需要类型的实例来调用。方法接收器必须有一个显式的名字，这个名字必须在方法中被使用。
//
//在接收器是指针时，方法可以改变接收器的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。
//
//在 Go 语言中，（接收器）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收器来建立。
//
//方法没有和定义的数据类型（结构体）混在一起，方法和数据是正交，而且数据和行为（方法）是相对独立的。

//指针方法与值方法
//有类型T，方法的接收器为(t T)时我们称为值接收器，该方法称为值方法；方法的接收器为(t *T)时我们称为指针接收器，该方法称为指针方法。
//
//如果想要方法改变接收器的数据，就在接收器的指针上定义该方法；否则，就在普通的值类型上定义方法。这是指针方法和值方法最大的区别

//无论你声明方法的接收器是指针接收器还是值接收器，Go都可以帮你隐式转换为正确的方法使用。
//
//但我们需要记住，值变量只拥有值方法集，而指针变量则同时拥有值方法集和指针方法集。

//怎么选择是指针接收器还是值接收器呢
//*何时使用值类型
//
//（1）如果接收器是一个 map，func 或者 chan，使用值类型（因为它们本身就是引用类型）。 （2）如果接收器是一个 slice，并且方法不执行 reslice 操作，也不重新分配内存给 slice，使用值类型。 （3）如果接收器是一个小的数组或者原生的值类型结构体类型(比如 time.Time 类型)，而且没有可修改的字段和指针，又或者接收器是一个简单地基本类型像是 int 和 string，使用值类型就好了。
//
//值类型的接收器可以减少一定数量的内存垃圾生成，值类型接收器一般会在栈上分配到内存（但也不一定），在没搞明白代码想干什么之前，别为这个原因而选择值类型接收器。
//
//*何时使用指针类型
//
//（1）如果方法需要修改接收器里的数据，则接收器必须是指针类型。 （2）如果接收器是一个包含了 sync.Mutex 或者类似同步字段的结构体，接收器必须是指针，这样可以避免拷贝。 （3）如果接收器是一个大的结构体或者数组，那么指针类型接收器更有效率。 （4）如果接收器是一个结构体，数组或者 slice，它们中任意一个元素是指针类型而且可能被修改，建议使用指针类型接收器，这样会增加程序的可读性。
//
//最后如果实在还是不知道该使用哪种接收器，那么记住使用指针接收器是最靠谱的。

//匿名类型的方法提升
//当一个匿名类型被嵌入在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型继承了这些方法：将父类型放在子类型中来实现亚型。
// 这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果。
//
//当我们嵌入一个匿名类型，这个类型的方法就变成了外部类型的方法，但是当它的方法被调用时，方法的接收器是内部类型(嵌入的匿名类型)，而非外部类型。
type People struct {
	Age    int
	gender string
	Name   string
}

//外部类型
type OtherPeople struct {
	//匿名类型
	People
}

func (p People) PeInfo() {
	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.gender)
}

//因此嵌入类型的名字充当着字段名，同时嵌入类型作为内部类型存在，我们可以使用下面的调用方法：
OtherPeople.People.PeInfo()

//这些字段和方法也同样被提升到了外部类型，我们可以直接访问：
OtherPeople.PeInfo()

//规则一：如果S包含嵌入字段T，则S和*S的方法集都包括具有接收方T的提升方法。*S的方法集还包括具有接收方*T的提升方法。
//
//规则二：如果S包含嵌入字段*T，则S和*S的方法集都包括具有接收器T或*T的提升方法。

import (
"fmt"
"reflect"
)

type People struct {
	Age    int
	gender string
	Name   string
}

type OtherPeople struct {
	People
}

type NewPeople People

func (p *NewPeople) PeName(pname string) {
	fmt.Println("pold name:", p.Name)
	p.Name = pname
	fmt.Println("pnew name:", p.Name)
}

func (p NewPeople) PeInfo() {
	fmt.Println("NewPeople ", p.Name, ": ", p.Age, "岁, 性别:", p.gender)
}

func (p *People) PeName(pname string) {
	fmt.Println("old name:", p.Name)
	p.Name = pname
	fmt.Println("new name:", p.Name)
}

func (p People) PeInfo() {
	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.gender)
}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("%T\n", a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(i, ":", m.Name, m.Type)
	}
}

func main() {
	p := OtherPeople{People{26, "Male", "张三"}}
	p.PeInfo()
	p.PeName("Joke")

	methodSet(p) // T方法提升

	methodSet(&p) // *T和T方法提升

	pp := NewPeople{42, "Male", "李四"}
	pp.PeInfo()
	pp.PeName("Haw")

	methodSet(&pp)
}

//程序输出：
//People  张三 :  26 岁, 性别: Male
//old name: 张三
//new name: Joke
//main.OtherPeople
//0 : PeInfo func(main.OtherPeople)
//*main.OtherPeople
//0 : PeInfo func(*main.OtherPeople)
//1 : PeName func(*main.OtherPeople, string)
//NewPeople  李四 :  42 岁, 性别: Male
//pold name: 李四
//pnew name: Haw
//*main.NewPeople
//0 : PeInfo func(*main.NewPeople)
//1 : PeName func(*main.NewPeople, string)

//虽然P 只有一个方法：PeInfo func(main.OtherPeople)，但我们依然可以调用p.PeName("Joke")。
//
//这里Go自动转为(&p).PeName("Joke")，其调用后结果让我们以为p有两个方法，其实这里p只有一个方法。