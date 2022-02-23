package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
    AccessKeyID string
    SecretAccessKey string
    MyRegion string
	Bucket string
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func Getenv() Config {
	LoadEnv()
	var Bucket string
	if GetEnvWithKey("ENVIRONMENT") ==  "production" {
		Bucket = GetEnvWithKey("AWS_BUCKET_PUBLIC")
	} else {
		Bucket = GetEnvWithKey("AWS_BUCKET_PUBLIC_TEST")
	}
	m := Config { GetEnvWithKey("AWS_ACCESS_KEY_ID"), GetEnvWithKey("AWS_SECRET_ACCESS_KEY"), GetEnvWithKey("AWS_REGION"), Bucket }
	return m
}
