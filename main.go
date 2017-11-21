package main

import (
	"github.com/pineapple-chicken/cloudgo-io/server"
)

func main() {
	s := server.NewServer()
	s.Run(":8080")
}
