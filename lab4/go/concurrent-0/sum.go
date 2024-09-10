package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Pathsum struct{
	_sum int
	path string
}

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string, ch chan Pathsum){
	data, err := readFile(filePath)
	if err != nil {
		
	}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	pathsum := Pathsum{_sum, filePath}
	ch <- pathsum
	
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	ch := make(chan Pathsum)
	var totalSum int
	sums := make(map[int][]string)
	for _, path := range os.Args[1:] {
		go sum(path, ch)
	}

	for _, _ = range os.Args[1:]{
		pathsum := <- ch
		totalSum += pathsum._sum
		sums[pathsum._sum] = append(sums[pathsum._sum], pathsum.path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}
