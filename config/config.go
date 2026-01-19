package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Conf struct {
	LogLevel          string `mapstructure:"LOG_LEVEL"`
	WebServerPort     int    `mapstructure:"WEB_SERVER_PORT"`
	HttpClientTimeout int    `mapstructure:"HTTP_CLIENT_TIMEOUT_MS"`
	ViaCepApiBaseUrl  string `mapstructure:"VIACEP_API_BASE_URL"`
	WeatherApiBaseUrl string `mapstructure:"WEATHER_API_BASE_URL"`
	WeatherApiKey     string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*Conf, error) {
	var c Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// If there is no config file, proceed and rely on environment variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok && !os.IsNotExist(err) {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	// Allow PORT environment variable (common in Cloud Run) to override
	if portStr, ok := os.LookupEnv("PORT"); ok && portStr != "" {
		// viper will parse int from env if set, but ensure it's used
		c.WebServerPort = viper.GetInt("PORT")
	}

	// Fallbacks to environment variables if viper didn't load them
	if c.ViaCepApiBaseUrl == "" {
		c.ViaCepApiBaseUrl = os.Getenv("VIACEP_API_BASE_URL")
	}

	if c.WeatherApiBaseUrl == "" {
		c.WeatherApiBaseUrl = os.Getenv("WEATHER_API_BASE_URL")
	}

	if c.WeatherApiKey == "" {
		c.WeatherApiKey = os.Getenv("WEATHER_API_KEY")
	}

	if c.HttpClientTimeout == 0 {
		if t := os.Getenv("HTTP_CLIENT_TIMEOUT_MS"); t != "" {
			if v, err := strconv.Atoi(t); err == nil {
				c.HttpClientTimeout = v
			}
		}
	}

	// If no port configured, error to avoid starting with port 0
	if c.WebServerPort == 0 {
		return nil, errors.New("web server port not configured (set PORT or WEB_SERVER_PORT)")
	}

	return &c, nil
}
