package client

import (
	"fmt"
	"log"
	"net"
	"time"
)

// NewBroadcaster creates a new UDP multicast connection on which to broadcast
func NewBroadcaster(address string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error connecting to " + address)
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting to " + address)
		return nil, err
	}

	return conn, nil

}

func ConnectToServer(address string) {
	conn, err := NewBroadcaster(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter text: ")
	var input string
	for {
		fmt.Scanln(&input)
		conn.Write([]byte(input))
		time.Sleep(1 * time.Second)
	}
}
