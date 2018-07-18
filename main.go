package main

import (
	"github.com/tnngo/ring/server"
	"fmt"
)

func main() {
	fmt.Println("how are you?")
	server.ListenTCP()
}
