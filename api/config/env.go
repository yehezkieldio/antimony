package config

import "github.com/spf13/viper"

type Configuration struct {
	LivekitAPIKey    string `mapstructure:"LIVEKIT_API_KEY"`
	LivekitAPISecret string `mapstructure:"LIVEKIT_API_SECRET"`
	LivekitHostURL   string `mapstructure:"LIVEKIT_HOST_URL"`

}

func LoadConfiguration(path string) (Configuration, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	var configuration Configuration
	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}