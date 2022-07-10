package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uchupx/golang-absensi/pkg/database"
	"github.com/uchupx/golang-absensi/pkg/redis"
)

type Config struct {
	AppName    string
	AppAddress string
	RSAKeyPath string

	DBConfig    database.MariaDBConfig
	Environment EnvironmentConfig
	RedisConfig redis.RedisConfig
}

var config *Config

func Init() {
	err := godotenv.Load("./conf/.env")
	if err != nil {
		log.Printf("[Init] error on loading env from file: %+v", err)
	}

	config = &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppAddress: os.Getenv("APP_ADDRESS"),
		RSAKeyPath: os.Getenv("RSA_KEY_PATH"),
		DBConfig: database.MariaDBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Host:     os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_ADDRESS"),
			DBName:   os.Getenv("DB_NAME"),
		},
		RedisConfig: redis.RedisConfig{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	env, ok := checkEnvironment(os.Getenv("ENVIRONMENT"))
	if !ok {
		log.Panicf("[Init] invalid environment, should be \"prod\" or \"dev\", found: %+v", os.Getenv("ENVIRONMENT"))
	}
	config.Environment = env

	if config.AppName == "" {
		log.Panicf("[Init] app name cannot be empty")
	}

	if config.AppAddress == "" {
		log.Panicf("[Init] app address cannot be empty")
	}

	if config.DBConfig.DBName == "" || config.DBConfig.Port == "" {
		log.Panicf("[Init] db name or address cannot be empty")
	}
}

func Get() *Config {
	return config
}
