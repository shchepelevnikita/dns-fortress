package udp

import (
	"fmt"
	"net"
)

func StartUDPServer() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
		Zone: "",
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Printf("Something went wrong when trying to start a UDP server: %v\n", err)
		return
	}

	fmt.Printf("Started a UDP server on %s\n", addr)

	for {
		var buf [512]byte

		_, addr, err := conn.ReadFromUDP(buf[0:])

		if err != nil {
			fmt.Printf("Something went wrong when trying to read from UDP server: %v\n", err)
			return
		}

		_, err = conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)

		if err != nil {
			fmt.Printf("Something went wrong when trying to send response over UDP server: %v\n", err)
			return
		}
	}
}
