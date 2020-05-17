package main

import (
	"fmt"
	"os"

	"github.com/jishnu001/Network/client"
	"github.com/jishnu001/Network/server"
)

func showUsage() {
	fmt.Print("Usage: ./Network --client serverAddress port\n")
	fmt.Print("Usage: ./Network --server serverAddress port\n")

}

func main() {

	if len(os.Args) < 4 {
		fmt.Print("Error, missing arguments\n")
		showUsage()
		return
	}
	mode := os.Args[1]
	address := os.Args[2]
	port := os.Args[3]

	if mode == "--client" {
		client.ConnectToServer(address + ":" + port)
	} else if mode == "--server" {
		server.StartServer(address + ":" + port)
	} else {
		fmt.Println("Unknown arguments")
		showUsage()
	}

}
