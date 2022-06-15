package main

import (
	"log"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	log.Println("Starting proxy server on port 443")
	if err := server.ListenAndServe("tcp", "0.0.0.0:443"); err != nil {
		panic(err)
	}
}
