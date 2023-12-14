package main

import (
	"github.com/gavotte25/blockchain_lab1/client"
	"github.com/gavotte25/blockchain_lab1/server"
)

func main() {
	server.Start()
	client.Start()
}
