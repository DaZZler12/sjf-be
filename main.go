package main

import (
	"fmt"

	"github.com/DaZZler12/sjf-be/cmd/sjf"
)

func main() {
	// get the server engine and start the server
	// handle the server shutdown
	fmt.Println("Hyee G'Day!")
	sjf.StartServer()
}
