package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/FishGoddess/cachego"
)

var (
	counter int64 = 0
)

func onMissed(ctx context.Context) (data interface{}, err error) {
	time.Sleep(time.Second) // Simulate blocking operations.
	return atomic.AddInt64(&counter, 1), nil
}

func main() {
	cache := cachego.NewCache()

	for i := 0; i < 10; i++ {
		value, err := cache.Get("key-reload", cachego.WithOpOnMissed(onMissed), cachego.WithOpTTL(time.Second))
		fmt.Println(value, err)
	}

	fmt.Println("=============================")

	for i := 0; i < 10; i++ {
		value, err := cache.Get("key-not-reload", cachego.WithOpOnMissed(onMissed), cachego.WithOpTTL(time.Second), cachego.WithOpDisableReload())
		fmt.Println(value, err)
	}
}
