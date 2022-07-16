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
		r.StatusCode = 500
		return r
	}

	defer resp.Body.Close()

	r.StatusCode = resp.StatusCode

	if r.StatusCode == 302 {
		r.Ip = resp.Header.Get("Location")
		return r
	} else if r.StatusCode == 200 {
		r.Header = parseBody(resp.Body)
	}

	r.Ip = getIP(url)
	return r
}

func getIP(url string) string {
	ip, err := net.ResolveIPAddr("ip4", parseUrl(url))
	if err != nil {
		return ""
	}
	return ip.String()
}
