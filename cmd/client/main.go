package main

import (
	"crypto/tls"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.DefaultClient
	client.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	for idx := 0; idx < 3; idx++ {
		response, err := client.Get("http://localhost:8088/")
		if err != nil {
			panic(response)
		}
	}
}
