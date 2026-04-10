package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func ListFiles() {
	files, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func MakeFIle() (*os.File, error) {
	file, err := os.Create("./data/output1.txt") // Assuming the directory already exists
	if err != nil {
		log.Fatal("Error creating the file")
	}
	return file, err
}

func WritingAFile() {
	file, _ := MakeFIle()

	defer file.Close() // This'd close the file after the wiritingAfile func is completed

	w := bufio.NewWriter(file)
	fmt.Fprint(w, "Hello\n My name is the name \n and I hate my college to the \n fullest")
	w.Flush()

	fmt.Println()
}

func CountingLines(filePath string) (int, error) {
	file, err := os.Open(filePath)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, err
}

func CountingLinesViaChannel(filepathChannel chan string, Wg *sync.WaitGroup) {

	defer Wg.Done()
	for filePath := range filepathChannel { // runs until channel gives input
		file, err := os.Open(filePath) // Opens up the file
		if err != nil {
			fmt.Println("Error : ", err)
			continue
		}

		defer file.Close()

		// to count the lines
		scanner := bufio.NewScanner(file)
		lineCount := 0

		for scanner.Scan() {
			lineCount++
		}
		fmt.Printf("Nums of line in file %v is %v\n", filePath, lineCount)
	}
}
