package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/maximekuhn/brainfuck/internal/webapp"
)

//go:embed banner.txt
var banner string

func main() {
	fmt.Println(banner)
	server := webapp.NewServer()
	log.Fatal(server.Run())
}
