package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/url"
	"strings"
)

func parseUrl(site string) string {
	address, err := url.Parse(site)
	if err != nil {
		panic(err)
	}
	hostname := address.Hostname()
	if strings.Contains(hostname, "www.") {
		hostname = strings.TrimPrefix(address.Hostname(), "www.")
	}
	return hostname
}

func parseBody(b io.ReadCloser) string {
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		panic(err)
	}
	title := doc.Find("title").Text()
	return title
}
