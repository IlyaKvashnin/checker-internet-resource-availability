package main

import (
	"sync"
)

const FilePath = "./files/rkn-70000.txt"

func main() {
	wg := sync.WaitGroup{}
	file := readFile(FilePath)
	for _, elem := range file {
		wg.Add(1)
		go checkUrl(&wg, elem)
	}
	wg.Wait()
}
