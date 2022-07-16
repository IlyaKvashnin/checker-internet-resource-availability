package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) []string {
	var data []string
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func logToFile() {
	f, err := os.OpenFile("./files/result.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		val, ok := <-c
		if !ok {
			break
		} else {
			if _, err = f.WriteString(fmt.Sprintln(val)); err != nil {
				fmt.Println(err)
			}
		}
	}
}
