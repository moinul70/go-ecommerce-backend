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
type DbConfig struct {
	Driver   string
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    Sslmode   string

    MaxOpen  int
    MaxIdle  int
    LifeTime int
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
func GetDbConfig() DbConfig{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		os.Exit(1)
	}

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		log.Println("error: driver is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Println("error: DB_PORT is required")
		os.Exit(1)
	}
	
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Println("error: user name is required")
		os.Exit(1)
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Println("error: password is required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Println("error: db name is required")
		os.Exit(1)
	}
	dbHost := os.Getenv("DB_HOST")
	if dbName == "" {
		log.Println("error: dbHost is required")
		os.Exit(1)
	}
	sslmode := os.Getenv("SSL_MODE")
	if sslmode == "" {
		log.Println("error: sslmode is required")
		os.Exit(1)
	}
	return DbConfig{
	Driver:   driver,
    Host:     dbHost,
    Port:     dbPort,
    User:     user,
    Password: password,
    DBName:   dbName,
    Sslmode:   sslmode,
    MaxOpen:  10,
    MaxIdle:  10,
    LifeTime: 30,
}
}
