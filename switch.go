package main

//import "fmt"

import (
	"fmt"
	"time"
)

//func main() {
//
//	//fallthrough强制执行后面的case代码，fallthrough不会判断下一条case的expr结果是否为true。
//	switch a := 3; {
//	case a == 1:
//		fmt.Println("The integer was == 1")
//		fallthrough
//	case a == 2:
//		fmt.Println("The integer was == 2")
//	case a == 3:
//		fmt.Println("The integer was == 3")
//		fallthrough
//	case a == 4:
//		fmt.Println("The integer was == 4")
//	case a == 6:
//		fmt.Println("The integer was == 4")
//	case a == 5:
//		fmt.Println("The integer was == 6")
//		fallthrough
//	default:
//		fmt.Println("default case")
//	}
//}

//如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
//如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
//如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行。

//select中的case语句必须是一个channel操作
//select中的default子句总是可运行的

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	case <-time.After(time.Second * 3): //超时退出
		fmt.Println("request time out")
		//default:
		//	fmt.Println("default")
	}

}
