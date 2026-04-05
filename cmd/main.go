package main

import (
	"bufio"
	"fmt"
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
	file, err := os.Create("./data/output1.txt")
	if err != nil {
		log.Fatal("Error creating the file")
	}

	// defer file.Close() #TODO: Do not close now will close when it's work is over in some other function

	return file, err
}

func writingAFile() {
	file, _ := MakeFIle()

	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world! ")
	w.Flush()

	fmt.Println()
}
