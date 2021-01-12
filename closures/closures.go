package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	netxInt := intSeq()
	fmt.Println(netxInt())
	fmt.Println(netxInt())
	fmt.Println(netxInt())

	newInts := intSeq()
	fmt.Println(newInts())
}