package main

import (
	"fmt"
	server "pttep-vr-api/server"
)

func main() {
	fmt.Println("Starting server")
	server.Run()
	fmt.Println("Server stopped")
}
