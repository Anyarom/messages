package config

import "github.com/spf13/viper"

type (
	Config struct {
		PgConfig  PgConfig
		RbtConfig RbtConfig
	}

	PgConfig struct {
		Host     string
		Port     int64
		Database string
		User     string
		Password string
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

	// для pg
	config.PgConfig = PgConfig{
		Host:     viper.GetString("pg.host"),
		Port:     viper.GetInt64("pg.port"),
		Database: viper.GetString("pg.database"),
		User:     viper.GetString("pg.user"),
		Password: viper.GetString("pg.password"),
	}

	// для rabbit
	config.RbtConfig = RbtConfig{
		Username: viper.GetString("rabbit.username"),
		Pass:     viper.GetString("rabbit.pass"),
		Host:     viper.GetString("rabbit.host"),
		Port:     viper.GetString("rabbit.port"),
	}

	return config, nil
}
