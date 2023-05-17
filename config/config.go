package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DatabaseUrl   string `mapstructure:"DATABASE_URL"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
}

var Env *Config

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Env)
	return
}
