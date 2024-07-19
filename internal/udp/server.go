package udp

import (
	"bytes"
	"github.com/spf13/viper"
	"io"
	"log"
	"net"
	"net/http"
)

func StartUDPServer() {

	ip := viper.GetString("server.ip")
	port := viper.GetInt("server.port")

	dnsServers := viper.GetStringMapString("dns-servers")

	addr := &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}

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

		for _, dnsServerUrl := range dnsServers {
			req, err := http.NewRequest("GET", dnsServerUrl+""+domain, nil)

			if err != nil {
				log.Printf("Something went wrong when trying to create GET request for resolving %s with %s: %v\n", domain, dnsServerUrl, err)
				continue
			}

			req.Header.Add("Accept", "application/dns-json")

			res, err := client.Do(req)

			if err != nil {
				log.Printf("Something went wrong when trying to forward domain %s to %s: %v\n", domain, dnsServerUrl, err)
				continue
			}

			resBody, err := io.ReadAll(res.Body)

			if err != nil {
				log.Printf("Something went wrong when trying to read response body: %v\n", err)
				continue
			}

			log.Printf("Response body: %s\nResponse code: %d\n", resBody, res.StatusCode)
			break
		}
	}
}
