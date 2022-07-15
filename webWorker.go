package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

func checkUrl(wg *sync.WaitGroup, url string) {
	r := response{url, "", 0, ""}

	var client = &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	//fmt.Println("Проверяем адрес ", url)
	resp, err := client.Get(url)

	if err != nil {
		r.statusCode = 500
		fmt.Println(r.ToString())
		//fmt.Printf("Ошибка соединения. %s\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		//fmt.Printf("Ошибка. http-статус: %v\n", resp.StatusCode)
		content, _ := ioutil.ReadAll(resp.Body)
		r.header = parseBody(string(content))
	}
	r.statusCode = resp.StatusCode
	r.ip = getIP(url)
	fmt.Println(r.ToString())
	//fmt.Printf("Онлайн. http-статус: %d\n", resp.StatusCode)
	wg.Done()
}

func getIP(url string) string {
	ip, err := net.ResolveIPAddr("ip4", url)
	if err != nil {
		defer recovery()
		panic(err)
	}
	return ip.String()
}
