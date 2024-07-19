package main

import (
	"github.com/dns-fortress-resolver/internal/udp"
	"github.com/spf13/viper"
	"log"
	"os"
)

func initConfig() {
	viper.AddConfigPath("config")

	if os.Getenv("ENV") == "PRODUCTION" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("local")
	}

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Couldn't read the configuration file", err)
	}
}

func main() {
	initConfig()

	udp.StartUDPServer()
}
