package testing

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestA(t *testing.T)  {
	// 创建一个3个元素缓冲大小的整型通道
	ch := make(chan int, 3)
	// 查看当前通道的大小
	fmt.Println(len(ch))

	go func() {
		for {
			c := <- ch
			fmt.Println( "------" , c)
		}
	}()
	ceshi(ch)

	// 查看当前通道的大小
	fmt.Println(len(ch))
}

func ceshi(ch chan int)  {
	// 发送3个整型元素到通道
	for i := 0; i < 3; i++ {
		ch <- i
		// 阻塞1秒
		time.Sleep(time.Second * 1)
	}
}


type A struct {
	lock  sync.Mutex
}

type B struct {
	lock   sync.Mutex
}


func TestB(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 3) //最多允许两个并发同时执行
	a := &A{}
	b := &B{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int, a *A, b *B) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			fmt.Println(id, time.Now())
			if id % 2 == 1 {
				a.lock.Lock()
				fmt.Println("=================")
				time.Sleep(time.Second * 3)
				a.lock.Unlock()
			} else {
				//b.lock.Lock()
				fmt.Println("------------------")
				time.Sleep(time.Second * 1)
				//b.lock.Unlock()
			}
		}(i, a , b)
	}

	wg.Wait()
}