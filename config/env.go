package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) string {
	loadConfig()

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion for key %s", key)
	}

	return value
}

func loadConfig() {
	// Read ENV variables
	viper.AutomaticEnv()

	// Fallback to default .env - return if file not exists, use it if it exists
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		return
	} else {
		viper.SetConfigFile(".env")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
