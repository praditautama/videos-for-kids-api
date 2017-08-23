package main

import (
	"flag"
	"fmt"
)

func main() {
	serverPtr := flag.String("server", "prod", "a bool")
	flag.Parse()
    fmt.Println("Server\t:", *serverPtr)
}