package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/FishGoddess/cachego"
	"github.com/FishGoddess/cachego/pkg/task"
)

func main() {
	// We provide a task for you to do some loops.
	t := task.Task{
		Before: func(ctx context.Context) {
			fmt.Println("Before...")
		},
		Fn: func(ctx context.Context) {
			fmt.Println("Fn...")
		},
		After: func(ctx context.Context) {
			fmt.Println("After...")
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Run runs a task which is usually called in a new goroutine.
	// go t.Run(ctx, time.Second)
	t.Run(ctx, time.Second)

	// You can use it to update your cache. Try this:
	cache := cachego.NewCache()

	t = task.Task{
		Before: func(ctx context.Context) {
			cache.Set("key", "before")
		},
		Fn: func(ctx context.Context) {
			cache.Set("key", strconv.FormatInt(rand.Int63n(100), 10))
		},
		After: func(ctx context.Context) {
			cache.Set("key", "after")
		},
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go t.Run(ctx, time.Second)

	// Simulate user requests
	for i := 0; i < 22; i++ {
		fmt.Println(cache.Get("key"))
		time.Sleep(500 * time.Millisecond)
	}
}
