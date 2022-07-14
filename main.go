package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
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

func checkUrl(wg *sync.WaitGroup, url string) {
	var client = &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	fmt.Println("Проверяем адрес ", url)
	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("Ошибка соединения. %s\n", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("Ошибка. http-статус: %s\n", resp.StatusCode)
		return
	}
	fmt.Printf("Онлайн. http-статус: %d\n", resp.StatusCode)
	wg.Done()
}

/*func checkIp(url string) {
	ips, err := net.LookupIP(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}*/
