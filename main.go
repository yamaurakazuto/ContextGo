package main

import (
	"context"
	"fmt"
	"time"
)

func gofunc(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	task := make(chan int, 5)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-task:
				fmt.Println("get", i)
			}

		}
	}()
	time.Sleep(time.Second)
	for i := 0; 5 > i; i++ {
		task <- i
	}
	time.Sleep(1 * time.Second)
	cancel()
}

func longProcess(ctx context.Context, ch chan<- string) {
	select {
	case <-time.After(3 * time.Second):
		ch <- "done"
	case <-ctx.Done():
		fmt.Println("キャンセルしました")
	}
}

func foo(parent context.Context) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	ch := make(chan string)
	go longProcess(ctx, ch)
	for {
		select {
		case <-ch:
			fmt.Println("success")
			return
		case <-ctx.Done():
			fmt.Println("timeout in foo")
			return
		}

	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go func() { fmt.Println("別ゴルーチン") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("そして時は動き出す")
	foo(context.Background())
	gofunc(context.Background())
	time.Sleep(2 * time.Second)

}
