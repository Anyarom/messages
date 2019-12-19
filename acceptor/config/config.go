package config

import "github.com/spf13/viper"

type (
	Config struct {
		RbtConfig RbtConfig
	}

	RbtConfig struct {
		Username string
		Pass     string
		Host     string
		Port     string
	}
)

// инициализация структуры
func InitConfig(file string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(file)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return config, err
	}

	config.RbtConfig = RbtConfig{
		Username: viper.GetString("rabbit.username"),
		Pass:     viper.GetString("rabbit.pass"),
		Host:     viper.GetString("rabbit.host"),
		Port:     viper.GetString("rabbit.port"),
	}

	return config, nil
}
