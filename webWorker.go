package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
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
	r.statusCode = resp.StatusCode
	if r.statusCode == 302 {
		r.ip = resp.Header.Get("Location")
		fmt.Println(r.ToString())
		return
	}
	if r.statusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		title := doc.Find("title").Text()
		r.header = title
	}
	r.ip = getIP(url)
	fmt.Println(r.ToString())
	//fmt.Printf("Онлайн. http-статус: %d\n", resp.StatusCode)
	wg.Done()
}

func getIP(url string) string {
	ip, err := net.ResolveIPAddr("ip4", parseUrl(url))
	if err != nil {
		defer recovery()
		panic(err)
	}
	return ip.String()
}
