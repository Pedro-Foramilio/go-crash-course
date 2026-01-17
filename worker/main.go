package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs) // 5 clients
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ { // 3 workers
		//start each worker that listens to a job channel and publishes results to a result channel
		go worker(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		//publish jobs to the job channels listenned by the workers
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-result
	}
}
