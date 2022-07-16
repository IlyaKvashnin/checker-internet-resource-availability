package main

import (
	"sync"
)

const FilePath = "./files/test1.txt"

var c chan response

func main() {
	file := readFile(FilePath)
	wg := sync.WaitGroup{}

	go func() {
		logToFile()
		wg.Done()
	}()

	c = make(chan response, 10000)
	routines := make(chan struct{}, 10000)

	for _, v := range file {
		url := v
		routines <- struct{}{}
		wg.Add(1)
		go func(url string) {
			c <- checkUrl(url)
			<-routines
			wg.Done()
		}(url)
	}

	wg.Wait()
	close(c)
}
