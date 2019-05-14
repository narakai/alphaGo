package main

//Go实现面向对象的两个关键是struct和interface，结构代替类，因为Go语言不提供类，但提供了结构体或自定义类型，
// 方法可以被添加到结构体或自定义类型中。结构体之间可以嵌套，类似继承。而interface定义接口，实现多态性。
//
//封装（数据隐藏）：
//
//和别的 OO 语言有 4 个或更多的访问层次相比，Go 把它简化为了 2 层： 1）包范围内的：通过标识符首字母小写，对象 只在它所在的包内可见 2）可导出的：通过标识符首字母大写，对象 对所在包以外也可见类型只拥有自己所在包中定义的方法。
//
//继承：
//
//Go没有显式的继承，而是通过组合实现继承，内嵌一个（或多个）包含想要的行为（字段和方法）的结构体；多重继承可以通过内嵌多个结构体实现。
//
//多态：
//
//多态是运行时特性，而继承则是编译时特征，也就是说，继承关系在编译时就已经确定了，而多态则可以实现运行时的动态绑定。Go用接口实现多态，某个类型的实例可以赋给它所实现的任意接口类型的变量。类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。Go 接口不是 Java 和 C# 接口的变体，而且：接口间是不相关的，并且是大规模编程和可适应的演进型设计的关键。
//
//另外Go没有构造函数，如果一定要在初始化对象的时候进行一些工作的话，可以自行封装产生实例的方法。实例化的时候可以初始化属性值，如果没有指明则默认为系统默认值。加&符号和new的是指针对象，没有的则是值对象，在传递对象的时候要根据实际情况来决定是要传递指针还是值。