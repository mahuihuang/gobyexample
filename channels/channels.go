package main

import "fmt"

func main() {
	messages := make(chan int)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
