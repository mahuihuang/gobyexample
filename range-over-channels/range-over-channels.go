package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	for elem := range queue {
		fmt.Println(elem)
	}
}
