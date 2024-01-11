package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"MARIADB_HOST"`
	DBUserName     string `mapstructure:"MARIADB_USER"`
	DBUserPassword string `mapstructure:"MARIADB_PASSWORD"`
	DBName         string `mapstructure:"MARIADB_DB"`
	DBPort         string `mapstructure:"MARIADB_PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
