package config

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// GetFromEnv reads environment varialbes to get configuration for the app
func GetConfiguration() (Config, error) {
	// Load variables from .env
	err := godotenv.Load()
	appForwardWebpack, err := strconv.ParseBool("false")
	if err != nil {
		return Config{}, errors.New("Invalid app forward webpack option specified")
	}
	config := Config{
		AppPort:           "4433",
		AppSecret:         []byte("secret"),
		AppTheme:          "default",
		AppForwardWebpack: appForwardWebpack,
		AppUserName:       "marcetin",
		AppUserPassword:   []byte("$2y$12$JbJmIwUnnXaaiZCYB8pNm.gt/KMdQ4/alJTrWWLkOasd3iNXZRpEu"),
		DBName:            "DATABASE",
	}
	err = validator.New().Struct(config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
