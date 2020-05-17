package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	fmt.Println(c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "stop" {
			break
		}

		fmt.Println(temp)

		result := "Got that...\n"
		c.Write([]byte(string(result)))
	}
	fmt.Println("Closing.")
	c.Close()
}

func StartServer(address string) {

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error binding to " + address)
		return
	}
	defer conn.Close()

	conn.SetReadBuffer(8192)
	buffer := make([]byte, 8192)

	for {
		n, add, e := conn.ReadFromUDP(buffer)
		if e != nil {
			continue
		}
		fmt.Println(add)
		strValue := string(buffer[:n])
		strValue = strings.TrimSpace(strValue)
		if len(strValue) > 0 {
			fmt.Println(strValue)
		}
	}
}

/*conn, err := net.Listen("tcp", ":8080")
if err != nil {
	fmt.Println("Error")
	return
}
conn, err := net.ListenPacket("udp", "239.0.0.0:9999")
*/
