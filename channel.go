package main

import (
	"fmt"
	"time"
)

//Go 奉行通过通信来共享内存，而不是共享内存来通信。所以，channel 是goroutine之间互相通信的通道，goroutine之间可以通过它发消息和接收消息。
//
//channel是进程内的通信方式，因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。
//
//channel是类型相关的，一个channel只能传递（发送或接受 | send or receive）一种类型的值，这个类型需要在声明channel时指定。
//
//默认的，信道的存消息和取消息都是阻塞的 (叫做无缓冲的信道)
//
//使用make来建立一个通道：
//var channel chan int = make(chan int)
//// 或
//channel := make(chan int)
// 定义接收的channel
//receive_only := make (<-chan int)

// 定义发送的channel
//send_only := make (chan<- int)

// 可同时发送接收
//send_receive := make (chan int)

//无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。
func main() {
	//无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。
	//c := make(chan int) // 不使用带缓冲区的channel
	//有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道
	c := make(chan int, 10) // 使用带缓冲区的channel
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}

// 只能向chan里send数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {

		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

// 只能接收channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}
