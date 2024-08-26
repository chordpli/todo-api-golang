package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	RDB                    string        `mapstructure:"rdb"`
	PORT                   string        `mapstructure:"port"`
	Access_Token_Duration  time.Duration `mapstructure:"accessTokenDuration"`
	Refresh_Token_Duration time.Duration `mapstructure:"refreshTokenDuration"`
	SecretKeyHex           string        `mapstructure:"secretKeyHex"`
	PublicKeyHex           string        `mapstructure:"publicKeyHex"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path) // <- to work with Dockerfile setup
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
