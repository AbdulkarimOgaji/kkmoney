package util

import "github.com/spf13/viper"

type Config struct {
	ServerAddr      string `mapstructure:"SERVER_ADDR"`
	DbconnectionUri string `mapstructure:"DB_CONNECTION_URI"`
	DbName          string `mapstructure:"DB_NAME"`
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
