package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
		Cors struct {
			AllowOrigins     []string `json:"allowOrigins"`
			AllowMethods     []string `json:"allowMethods"`
			AllowHeaders     []string `json:"allowHeaders"`
			ExposeHeaders    []string `json:"exposeHeaders"`
			AllowCredentials bool     `json:"allowCredentials"`
			MaxAge           string   `json:"maxAge"`
		} `json:"cors"`
	} `json:"server"`
	Database struct {
		DataSourceName    string `json:"dataSourceName"`
		MaxIdleConns      int    `json:"maxIdleConns"`
		MaxOpenConns      int    `json:"maxOpenConns"`
		DefaultQueryLimit string `json:"defaultQueryLimit"`
		ConnMaxLifetime   string `json:"connMaxLifetime"`
	} `json:"database"`
	Minio struct {
		Endpoint            string `json:"endpoint"`
		AccessKey           string `json:"accessKey"`
		SecretKey           string `json:"secretKey"`
		Secure              bool   `json:"secure"`
		BucketName          string `json:"bucketName"`
		BucketRegion        string `json:"bucketRegion"`
		BucketObjectLocking bool   `json:"bucketObjectLocking"`
	} `json:"minio"`
}

var Conf *Config // Global configuration accessible throughout the application

func LoadConfig() {
	path := "config.json"
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	Conf = &Config{}
	if err = decoder.Decode(Conf); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	log.Println("Configuration loaded successfully")
}
