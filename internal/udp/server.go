package udp

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
)

func StartUDPServer() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
		Zone: "",
	}

	cloudflareURL := "https://1.1.1.1/dns-query?name="

	client := &http.Client{}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Printf("Something went wrong when trying to start a UDP server: %v\n", err)
		return
	}

	log.Printf("Started a UDP server on %s\n", addr)

	for {
		var buf [512]byte

		_, _, err := conn.ReadFromUDP(buf[0:])

		if err != nil {
			log.Printf("Something went wrong when trying to read from UDP server: %v\n", err)
			return
		}

		domain := string(bytes.Trim(buf[:], "\x00"))

		req, err := http.NewRequest("GET", cloudflareURL+""+domain, nil)

		if err != nil {
			log.Printf("Something went wrong when trying to create GET request for resolving %s with Cloudflare: %v\n", domain, err)
			return
		}

		req.Header.Add("Accept", "application/dns-json")

		res, err := client.Do(req)

		if err != nil {
			log.Printf("Something went wrong when trying to forward domain %s to Cloudflare: %v\n", domain, err)
			return
		}

		resBody, err := io.ReadAll(res.Body)

		if err != nil {
			log.Printf("Something went wrong when trying to read response body: %v\n", err)
			return
		}

		log.Printf("Response body: %s\nResponse code: %d\n", resBody, res.StatusCode)
	}
}
