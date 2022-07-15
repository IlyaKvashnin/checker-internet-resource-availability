package main

import (
	"fmt"
	"sync"
)

const FilePath = "./files/rkn-70000.txt"

func main() {
	wg := sync.WaitGroup{}
	file := readFile(FilePath)
	for idx, elem := range file {
		fmt.Println(idx)
		wg.Add(1)
		go checkUrl(&wg, elem)
	}
	wg.Wait()
}
