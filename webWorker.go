package main

import (
	"net"
	"net/http"
	"time"
)

func checkUrl(url string) string {
	r := response{url, "", 0, ""}

	var client = &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	resp, err := client.Get(url)
	if err != nil {
		r.statusCode = 500
		return r.ToString()
	}

	defer resp.Body.Close()

	r.statusCode = resp.StatusCode

	if r.statusCode == 302 {
		r.ip = resp.Header.Get("Location")
		return r.ToString()
	} else if r.statusCode == 200 {
		r.header = parseBody(resp.Body)
	}

	r.ip = getIP(url)
	return r.ToString()
}

func getIP(url string) string {
	ip, err := net.ResolveIPAddr("ip4", parseUrl(url))
	if err != nil {
		panic(err)
	}
	return ip.String()
}
