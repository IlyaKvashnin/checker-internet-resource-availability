package main

import (
	"bufio"
	"log"
	"os"
)

func readFile(path string) []string {
	var data []string
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
