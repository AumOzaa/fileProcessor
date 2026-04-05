package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"log"
	"os"
)

func main() {
	// Listing files
	fmt.Println("--------- Exisiting Files ---------")
	listFiles()
	fmt.Println("-------------------------------")

	fmt.Println("--------- Writing the file ---------")
	writingAFile()
	fmt.Println("-------------------------------")
}

func listFiles() {
	files, err := os.ReadDir(".")
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

	// defer file.Close() #TODO: Do not close now will close when it's work is over in some other function

	return file, err
}

func writingAFile() {
	file, _ := MakeFIle()

	defer file.Close() // This'd close the file after the wiritingAfile func is completed

	w := bufio.NewWriter(file)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world! ")
	w.Flush()

	fmt.Println()
}

func countingLines(filePath string) (int, error) {
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
