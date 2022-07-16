package main

import (
	"bufio"
	"encoding/json"
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
	f, err := os.OpenFile("./files/res.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()
	for {
		val, ok := <-c
		if !ok {
			break
		} else {
			result, e := json.MarshalIndent(val, "", " ")
			if e != nil {
				panic(e)
			}
			if _, err = f.WriteString(string(result[:])); err != nil {
				panic(err)
			}
		}
	}
}
