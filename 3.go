package main

import "fmt"

func main() {
	data := make(chan int)

	select {
	case number := <-data:
		fmt.Println("данные получены:", number)

	default:
		fmt.Println("нет данных")
	}
}