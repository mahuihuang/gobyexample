package main

import "fmt"

func varls() (int, int) {
	return 3, 7
}

func main() {
	a, b := varls()
	fmt.Println(a)
	fmt.Println(b)

	_, c := varls()
	fmt.Println(c)
}