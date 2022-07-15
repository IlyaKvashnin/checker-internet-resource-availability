package main

import "strconv"

type response struct {
	url        string
	ip         string
	statusCode int
	header     string
}

func (r *response) ToString() string {
	return r.url + " | " + r.ip + " | " + strconv.Itoa(r.statusCode) + " | " + r.header + "\n"
}
