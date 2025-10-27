package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Name     string
	Host     string
	PORT     int
	USER     string
	Password string
	SslMode  bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	SecretKey    string
	DbConnection DBConfig
}

var configInfo *Config

func throwErrorForConfig(envName string) {
	fmt.Println(envName, "IS Required")
	os.Exit(1)
}

func loadConfig() {
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
	//db relted config
	host := os.Getenv("HOST")
	if host == "" {
		throwErrorForConfig("HOST")
	}
	dbName := os.Getenv("NAME")
	if dbName == "" {
		throwErrorForConfig("DB Name")
	}
	dbUser := os.Getenv("USER")
	if dbUser == "" {
		throwErrorForConfig("DB USER")
	}
	dbPassword := os.Getenv("PASSWORD")
	if dbPassword == "" {
		throwErrorForConfig("DB PASSWORD")
	}
	dbPrt := os.Getenv("PORT")
	if dbPrt == "" {
		throwErrorForConfig("DB PASSWORD")
	}
	dbPort, err := strconv.ParseInt(dbPrt, 10, 64)
	if err != nil {
		throwErrorForConfig("DB POrt {cannt convert INT}")
	}
	SslModeDB := os.Getenv("SSL_MODE")
	sslModeBool := false
	if SslModeDB == "true" {
		sslModeBool = true
	}

	dbConfig := &DBConfig{
		Host:     host,
		Name:     dbName,
		Password: dbPassword,
		SslMode:  sslModeBool,
		PORT:     int(dbPort),
		USER:     dbUser,
	}

	port, err := strconv.ParseInt(ServicePORT, 10, 64)
	if err != nil {
		throwErrorForConfig("HTTP_PORT")
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")

	configInfo = &Config{
		Version:      ServiceVersion,
		ServiceName:  ServiceName,
		HttpPort:     int(port),
		SecretKey:    secretKey,
		DbConnection: *dbConfig,
	}
}

func GetConfig() *Config {
	if configInfo == nil {
		loadConfig()
	}
	return configInfo
}
