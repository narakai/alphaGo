package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

//通过go关键字让我们很容易启动一个协程，但难的是很好的管理和控制他们的运行。有几种方法我们可以根据场景使用：
//
//（1）使用sync.WaitGroup，它用于线程总同步，会等待一组线程集合完成，才会继续向下执行，这对监控所有子协程全部完成情况特别有用，但要控制某个协程就无能为力了；
//
//（2）使用channel来传递消息，一个协程来发送channel信号，另一个协程通过select来得到channel信息，这种方式可以满足协程之间的通信，来控制协程运行。但如果协程数量达到一定程度，就很难把控了；或者这两个协程还和其他协程也有类似通信，比如A与B，B与C，如果A发信号B退出了，C有可能等不到B的channel信号而被遗忘；
//
//（3）使用Context来传递消息，Context是层层传递机制，根节点完全控制了子节点，根节点（父节点）可以根据需要选择自动还是手动结束子节点。而每层节点所在的协程就可以根据信息来决定下一步的操作。

var logs *log.Logger

//func doClearn(ctx context.Context) {
//	// for 循环来每1秒work一下，判断ctx是否被取消了，如果是就退出
//	for {
//		time.Sleep(1 * time.Second)
//		select {
//		case <-ctx.Done():
//			logs.Println("doClearn:收到Cancel，做好收尾工作后马上退出。")
//			return
//		default:
//			logs.Println("doClearn:每隔1秒观察信号，继续观察...")
//		}
//	}
//}
//
//func doNothing(ctx context.Context) {
//	for {
//		time.Sleep(3 * time.Second)
//		select {
//		case <-ctx.Done():
//			logs.Println("doNothing:收到Cancel，但不退出......")
//
//			// 注释return可以观察到，ctx.Done()信号是可以一直接收到的，return不注释意味退出协程
//			//return
//		default:
//			logs.Println("doNothing:每隔3秒观察信号，一直运行")
//		}
//	}
//}
//
//func main() {
//	logs = log.New(os.Stdout, "", log.Ltime)
//
//	// 新建一个ctx
//	ctx, cancel := context.WithCancel(context.Background())
//
//	// 传递ctx
//	go doClearn(ctx)
//	go doNothing(ctx)
//
//	// 主程序阻塞20秒，留给协程来演示
//	time.Sleep(20 * time.Second)
//	logs.Println("cancel")
//
//	// 调用cancel：context.WithCancel 返回的CancelFunc
//	cancel()
//
//	// 发出cancel 命令后，主程序阻塞10秒，再看协程的运行情况
//	time.Sleep(10 * time.Second)
//}

//这里用Context嵌套控制3个协程，A，B，C。在主程序发出cancel信号后，每个协程都能接收根Context的Done()信号而退出。
func A1(ctx context.Context) int {
	ctx = context.WithValue(ctx, "AFunction", "Great")

	go B1(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("A Done")
		return -1
	}
	return 1
}

func B1(ctx context.Context) int {
	fmt.Println("A value in B:", ctx.Value("AFunction"))
	ctx = context.WithValue(ctx, "BFunction", 999)

	go C(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("B Done")
		return -2
	}
	return 2
}

func C(ctx context.Context) int {
	fmt.Println("B value in C:", ctx.Value("AFunction"))
	fmt.Println("B value in C:", ctx.Value("BFunction"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		fmt.Println("C Done")
		return -3
	}
	return 3
}

func main() {
	// 自动取消(定时取消)
	{
		timeout := 10 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		fmt.Println("A 执行完成，返回：", A1(ctx))
		select {
		case <-ctx.Done():
			fmt.Println("context Done")
			break
		}
	}
	time.Sleep(20 * time.Second)
}

//ContextMiddle是http服务中间件，统一读取通行cookie并使用ctx传递