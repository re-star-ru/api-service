package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"sync"
	"time"
)

var cfg Config
var once sync.Once

type Config struct {
	LelLel string   `mapstructure:"lel"`
	Kek    []string `mapstructure:"kek" hcl:"kek"`
}

// TODO: update remote config

func GetConfig() *Config {
	once.Do(func() {
		println("lel do get func")
		viper.SetDefault("HOST", "0.0.0.0")
		viper.SetDefault("PORT", "8080")
		viper.SetDefault("CONSUL_ADDR", "http://127.0.0.1:8500")
		viper.AutomaticEnv()

		if err := viper.AddRemoteProvider(
			"consul", viper.GetString("CONSUL_ADDR"), "/configs/api-service/api",
		); err != nil {
			log.Fatalln(err)
		}
		viper.SetConfigType("toml")

		if err := viper.ReadRemoteConfig(); err != nil {
			log.Fatalln(err)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalln(err)
		}
	})

	go updater()

	return &cfg
}

func updater() {
	for {
		time.Sleep(time.Second * 5)
		if err := viper.WatchRemoteConfig(); err != nil {
			log.Println(err)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalln(err)
		}

		log.Println("update config from consul")
		log.Println(viper.GetStringSlice("kek"))
		log.Println(cfg)
	}
}
