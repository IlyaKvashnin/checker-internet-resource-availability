package main

import (
	"net"
	"net/http"
	"time"
)

func checkUrl(url string) response {
	r := response{url, "", 0, ""}

	var client = &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	resp, err := client.Get(url)
	if err != nil {
		r.statusCode = 500
		return r
	}

	defer resp.Body.Close()

	r.statusCode = resp.StatusCode

	if r.statusCode == 302 {
		r.ip = resp.Header.Get("Location")
		return r
	} else if r.statusCode == 200 {
		r.header = parseBody(resp.Body)
	}
	r.ip = getIP(url)
	return r
}

func getIP(url string) string {
	ip, err := net.ResolveIPAddr("ip4", parseUrl(url))
	if err != nil {
		panic(err)
	}
	return ip.String()
}
