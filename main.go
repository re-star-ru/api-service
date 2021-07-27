package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"net/http"
)

func main() {
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
	vip.SetConfigType("json")

	if err := vip.ReadRemoteConfig(); err != nil {
		log.Fatalln(err)
	}

	log.Println("values:", vip.Get("lel"))

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello world"); err != nil {
			log.Println("error /:", err)
		}
	})

	var addr string
	if vip.GetString("ADDR") != "" {
		addr = vip.GetString("ADDR")
	} else {
		addr = fmt.Sprintf("%s:%s", vip.GetString("HOST"), vip.GetString("PORT"))
	}

	log.Println("Server listen at:", addr)
	log.Fatalln(http.ListenAndServe(addr, r))
}
