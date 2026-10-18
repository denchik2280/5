package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		result <- "ответ получен"
	}()

	select {
	case response := <-result:
		fmt.Println(response)

	case <-time.After(2 * time.Second):
		fmt.Println("время ожидания вышло!")
	}
}
