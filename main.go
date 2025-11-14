package main

import (
	"context"
	"fmt"
	"time"
)

func animal() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("cute")
		case <-ctx.Done():
			fmt.Println("キャンセル")
			return
		}
	}
}

func main() {
	animal()
}
