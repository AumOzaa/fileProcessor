package main

import (
	// "fmt"
	"sync"

	"github.com/AumOzaa/fileProcessor/functions"
)

func main() {
	// Listing files
	// fmt.Println("--------- Exisiting Files ---------")
	// functions.ListFiles()
	// fmt.Println("--------- Writing the file ---------")
	// functions.WritingAFile()
	// fmt.Println("--------- Reading lines from a file ---------")
	// numberOfLines, _ := functions.CountingLines("./data/output1.txt")
	// fmt.Println("Number of line(s) are ", numberOfLines)

	listOfFiles := []string{"./data/output1.txt", "./data/output2.txt", "./data/output3.txt"}

	workerCount := 2
	chann := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1) // 2. Tell the counter "We are waiting for 1 more worker"
		go func() {
			defer wg.Done() // 3. Tell the counter "I'm finished" when the function ends
			functions.CountingLinesViaChannel(chann)
		}()
	}
	for _, v := range listOfFiles {
		chann <- v
	}

	close(chann)
	wg.Wait()
}
