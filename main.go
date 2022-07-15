package main

import (
	"sync"
)

const FilePath = "./files/test1.txt"

func main() {
	wg := sync.WaitGroup{}
	file := readFile(FilePath)
	for _, elem := range file {
		wg.Add(2)
		go checkUrl(&wg, elem)
	}
	wg.Wait()
}
