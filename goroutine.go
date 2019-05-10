package main

import (
	"fmt"
	"time"
)

struct  G
{
uintptr stackguard0;// 用于栈保护，但可以设置为StackPreempt，用于实现抢占式调度
uintptr stackbase;  // 栈顶
Gobuf   sched;      // 执行上下文，G的暂停执行和恢复执行，都依靠它
uintptr stackguard; // 跟stackguard0一样，但它不会被设置为StackPreempt
uintptr stack0;     // 栈底
uintptr stacksize;  // 栈的大小
int16   status;     // G的六个状态
int64   goid;       // G的标识id
int8*   waitreason; // 当status==Gwaiting有用，等待的原因，可能是调用time.Sleep之类
G*  schedlink;      // 指向链表的下一个G
uintptr gopc;       // 创建此goroutine的Go语句的程序计数器PC，通过PC可以获得具体的函数和代码行数
};

struct P
{
Lock;       // plan9 C的扩展语法，相当于Lock lock;
int32   id;  // P的标识id
uint32  status;     // P的四个状态
P*  link;       // 指向链表的下一个P
M*  m;      // 它当前绑定的M，Pidle状态下，该值为nil
MCache* mcache; // 内存池
// Grunnable状态的G队列
uint32  runqhead;
uint32  runqtail;
G*  runq[256];
// Gdead状态的G链表（通过G的schedlink）
// gfreecnt是链表上节点的个数
G*  gfree;
int32   gfreecnt;
};

struct  M
{
G*  g0;     // M默认执行G
void    (*mstartfn)(void);  // OS线程执行的函数指针
G*  curg;       // 当前运行的G
P*  p;      // 当前关联的P，要是当前不执行G，可以为nil
P*  nextp;  // 即将要关联的P
int32   id; // M的标识id
M*  alllink;    // 加到allm，使其不被垃圾回收(GC)
M*  schedlink;  // 指向链表的下一个M
};

//在Go语言中，协程(goroutine)的使用很简单，直接在函数（代码块）前加上关键字 go 即可。
// go关键字就是用来创建一个协程(goroutine)的，后面的代码块就是这个协程(goroutine)需要执行的代码逻辑。

package main

import (
"fmt"
"time"
)

func main() {
	for i := 1; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	// 暂停一会，保证打印全部结束
	time.Sleep(1e9)
}
