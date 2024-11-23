package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
