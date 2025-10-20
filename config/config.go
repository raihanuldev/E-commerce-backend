package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	SecretKey   string
}

var configInfo Config

func throwErrorForConfig(envName string) {
	fmt.Println(envName, "IS Required")
	os.Exit(1)
}

func loadConfig() {
	// ✅ 2. .env file load করো (PRIMA!)
	godotenv.Load()

	ServiceName := os.Getenv("SERVICE_NAME")
	if ServiceName == "" {
		throwErrorForConfig("SERVICE_NAME")
	}

	ServiceVersion := os.Getenv("VERSION")
	if ServiceVersion == "" {
		throwErrorForConfig("VERSION")
	}

	ServicePORT := os.Getenv("HTTP_PORT")
	if ServicePORT == "" {
		throwErrorForConfig("HTTP_PORT")
	}

	port, err := strconv.ParseInt(ServicePORT, 10, 64)
	if err != nil {
		throwErrorForConfig("HTTP_PORT")
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")

	configInfo = Config{
		Version:     ServiceVersion,
		ServiceName: ServiceName,
		HttpPort:    int(port),
		SecretKey:   secretKey,
	}
}

func GetConfig() *Config {
	loadConfig()
	return &configInfo
}
