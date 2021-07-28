package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"srvapi/config"
)

func main() {
	cfg := config.GetConfig()

	log.Println("struct:", cfg)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello world"); err != nil {
			log.Println("error /:", err)
		}
	})

	var addr string

	if viper.GetString("ADDR") != "" {
		addr = viper.GetString("ADDR")
	} else {
		addr = fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
	}

	log.Println("Server listen at:", addr)
	log.Fatalln(http.ListenAndServe(addr, r))
}
