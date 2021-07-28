package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfg Config

type Config struct {
	LelLel string `mapstructure:"lel"`
	Kek    []string
}

func GetConfig() *Config {
	vip := viper.New()
	vip.SetDefault("HOST", "0.0.0.0")
	vip.SetDefault("PORT", "8080")
	vip.SetDefault("CONSUL_ADDR", "http://127.0.0.1:8500")
	vip.AutomaticEnv()

	if err := vip.AddRemoteProvider(
		"consul", vip.GetString("CONSUL_ADDR"), "/configs/api-service/api",
	); err != nil {
		log.Fatalln(err)
	}
	vip.SetConfigType("hcl")

	if err := vip.ReadRemoteConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := vip.Unmarshal(&cfg); err != nil {
		log.Fatalln(err)
	}

	return &cfg
}
