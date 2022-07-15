package main

import (
	"bufio"
	"fmt"
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

func recovery() {
	if msg := recover(); msg != nil {
		fmt.Println(msg)
	}
}
