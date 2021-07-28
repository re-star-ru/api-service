package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"log"
)

var cfg Config

type Config struct {
	LelLel string `mapstructure:"lel"`
	Kek    []string
}

// TODO: update remote config

func GetConfig() *Config {
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("CONSUL_ADDR", "http://127.0.0.1:8500")
	viper.AutomaticEnv()

	if err := viper.AddRemoteProvider(
		"consul", viper.GetString("CONSUL_ADDR"), "/configs/api-service/api",
	); err != nil {
		log.Fatalln(err)
	}
	viper.SetConfigType("hcl")

	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalln(err)
	}

	return &cfg
}
