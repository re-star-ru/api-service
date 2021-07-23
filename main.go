package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	cfg := viper.New()
	cfg.AddRemoteProvider("consul", consulAddr, consulPath)

	cfg := Config{}
	r := chi.NewRouter()

	viper.New()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello world"); err != nil {
			log.Println("error /:",err)
		}
	})

	log.Fatalln(http.ListenAndServe(cfg.host, r))
}

type Config struct {
	host string
}

type new() {
	v := viper
}