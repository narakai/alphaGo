package main

//Go语言包中的sync包提供了两种锁类型：sync.Mutex和sync.RWMutex，前者是互斥锁，后者是读写锁。

//建议：同一个互斥锁的成对锁定和解锁操作放在同一层次的代码块中。 使用锁的经典模式：
//var lck sync.Mutex
//func foo() {
//	lck.Lock()
//	defer lck.Unlock()
//	// ...
//}
//lck.Lock()会阻塞直到获取锁，然后利用defer语句在函数返回时自动释放锁。

import (
	"fmt"
	"sync"
	"time"
)

//当有锁释放时，才能进行lock动作，G0锁释放时，才有后续锁释放的可能，这里是G1抢到释放机会。
//func main() {
//	wg := sync.WaitGroup{}
//
//	var mutex sync.Mutex
//	fmt.Println("Locking  (G0)")
//	mutex.Lock()
//	fmt.Println("locked (G0)")
//	wg.Add(3)
//
//	for i := 1; i < 4; i++ {
//		go func(i int) {
//			fmt.Printf("Locking (G%d)\n", i)
//			mutex.Lock()
//			fmt.Printf("locked (G%d)\n", i)
//
//			time.Sleep(time.Second * 2)
//			mutex.Unlock()
//			fmt.Printf("unlocked (G%d)\n", i)
//			wg.Done()
//		}(i)
//	}
//
//	time.Sleep(time.Second * 5)
//	fmt.Println("ready unlock (G0)")
//	mutex.Unlock()
//	fmt.Println("unlocked (G0)")
//
//	wg.Wait()
//}

//WaitGroup，它用于线程同步，WaitGroup等待一组线程集合完成，才会继续向下执行。
// 主线程(goroutine)调用Add来设置等待的线程(goroutine)数量。 然后每个线程(goroutine)运行，并在完成后调用Done

//Mutex也可以作为struct的一部分，这样这个struct就会防止被多线程更改数据。
type Book struct {
	BookName string
	L        *sync.Mutex
}

func (bk *Book) SetName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		bk.L.Unlock()
		wg.Done()
	}()

	bk.L.Lock()
	fmt.Println("Lock set name:", name)
	time.Sleep(1 * time.Second)
	bk.BookName = name
}

func main() {
	bk := Book{}
	bk.L = new(sync.Mutex)
	wg := &sync.WaitGroup{}
	books := []string{"《三国演义》", "《道德经》", "《西游记》"}
	for _, book := range books {
		wg.Add(1)
		go bk.SetName(wg, book)
	}

	wg.Wait()
}

//sync.Once.Do(f func())能保证once只执行一次,这个sync.Once块只会执行一次。

//随着Go1.9的发布，有了一个新的特性，那就是sync.map，它是原生支持并发安全的map。虽然说普通map并不是线程安全（或者说并发安全），
// 但一般情况下我们还是使用它，因为这足够了；只有在涉及到线程安全，再考虑sync.map。
//
//但由于sync.Map的读取并不是类型安全的，所以我们在使用Load读取数据的时候我们需要做类型转换。