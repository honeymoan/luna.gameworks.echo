package main

import (
	"fmt"

	"go-micro.dev/v5"
)

func main() {
	service := micro.New("luna.gameworks.echo")
	service.Run()
	fmt.Println("Hello, World!")
}
