package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBName     string `mapstructure:"DB_NAME"`
	DBAddress  string `mapstructure:"DB_ADDRESS"`
	DBPort     int32  `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    int32  `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetDBConnectionURI(config *Config) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBAddress,
		config.DBPort,
		config.DBName,
	)
}

func GetServerURI(config *Config) string {
	return fmt.Sprintf(
		"%s:%d",
		config.ServerAddress,
		config.ServerPort,
	)
}
