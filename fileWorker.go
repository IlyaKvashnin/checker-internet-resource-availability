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

func log_to_file(s string) {
	// Сохраняет сообщения в файл
	f, err := os.OpenFile("./files/result.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if _, err = f.WriteString(fmt.Sprintln(s)); err != nil {
		fmt.Println(err)
	}
}
