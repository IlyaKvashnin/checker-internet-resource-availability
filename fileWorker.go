package main

import (
	"bufio"
	"encoding/json"
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
	f, err := os.OpenFile("./files/res.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
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

			r, _ := json.Marshal(val)
			fmt.Printf("%s", r)
			if _, err = f.WriteString(string(r)); err != nil {
				panic(err)
			}
			//data, fatal := json.Marshal(val)
			//if fatal != nil {
			//	panic(err)
			//} else {
			//	if _, err = f.WriteString(string(data)); err != nil {
			//		panic(err)
			//	}
			//}

		}
	}
}
