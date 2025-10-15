// libs/config/loader.go
package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func Load(path string) (*Config, error) {

	// Loading .env file
	// Checking in which mode the program is launching
	// Depends on the "ENV" variable
	// Takes values: "development", "local", "production"
	env := os.Getenv("APP_ENV")
	if env == "local" || env == "" {
		err := godotenv.Load(path + ".env")
		if err != nil {
			log.Fatalf("INFO: .env file not found at %s\n", path)
		}
	}

	// Initialize viper
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// Set up automatic parsing .env file in yaml structure
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Reading config file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to load cfg file: %w", err)
	}

	// Unmarshal all config
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cfg file: %w", err)
	}

	// Validating config values
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate cfg file: %w", err)
	}

	return &cfg, nil
}
