package main

import (
	"log"
	"net/url"
	"strings"
)

func parseUrl(site string) string {
	address, err := url.Parse(site)
	if err != nil {
		log.Fatal(err)
	}
	hostname := address.Hostname()
	if strings.Contains(hostname, "www.") {
		hostname = strings.TrimPrefix(address.Hostname(), "www.")
	}
	return hostname
}
