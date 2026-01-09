package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version         string
	HttpPort        int
	ApplicationName string
}
var configuration Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Println("error: version is required")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 64)
	if err != nil {
		log.Fatal("Bad HTTP_PORT:", err)
	}
	if port < 1 || port > 65535 {
		log.Fatal("HTTP_PORT must be between 1 and 65535")
	}
	httpPort := int(port)
	applicationName := os.Getenv("APPLICATION_NAME")
	if applicationName == "" {
		log.Println("error: application name is required")
		os.Exit(1)
	}
	configuration = Config{
		Version:         version,
		HttpPort:        httpPort,
		ApplicationName: applicationName,
	}

}
func GetEnv() Config{
	loadConfig()
	return configuration

}
