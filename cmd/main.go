package main

import (
	"fmt"
	// "sync"
	"time"
	// "github.com/AumOzaa/fileProcessor/functions"
)

func main() {
	// listOfFiles := []string{"./data/output1.txt", "./data/output2.txt", "./data/output3.txt"}
	//
	// workerCount := 2
	// chann := make(chan string)
	//
	// var wg sync.WaitGroup
	//
	// for i := 0; i < workerCount; i++ {
	// 	wg.Add(1) // 2. Tell the counter "We are waiting for 1 more worker"
	// 	go func() {
	// 		defer wg.Done() // 3. Tell the counter "I'm finished" when the function ends
	// 		functions.CountingLinesViaChannel(chann)
	// 	}()
	// }
	//
	// for _, v := range listOfFiles {
	// 	chann <- v
	// }
	//
	// close(chann)
	// wg.Wait()

	const numOfJobs = 1
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	fmt.Println("For loop is over now") // so we hired workers first and then now we'd give them some work

	for j := 1; j <= 4; j++ { // givinig work to workers
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numOfJobs; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker ", id, " started job ", j)
		time.Sleep(time.Second)

		fmt.Println("Worker ", id, " finsihed his job ", j)
		results <- j * 2
	}
}
