package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("wokrer", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 3 个 goroutine 读取数据
	for w := 1; w < 3; w++ {
		go worker(w, jobs, results)
	}

	// 写入 5 条数据
	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 读取数据
	for a := 1; a < numJobs; a++ {
		<-results
	}
}
