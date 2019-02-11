package config

import "github.com/spf13/viper"

type Constants struct {
	PORT string
}

type Config struct {
	Constants
}

func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants

	if err != nil {
		return &config, err
	}

	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("config") // Configuration fileName without the .TOML or .YAML or .json extension
	viper.AddConfigPath(".")      // Search the root directory for the configuration file

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return Constants{}, err
	}

	viper.SetDefault("PORT", "3000")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
