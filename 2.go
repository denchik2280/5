package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer completion.Done()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("работаем")

			case <-done:
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	done <- struct{}{}
	wg.Wait()

	fmt.Println("все")
}
