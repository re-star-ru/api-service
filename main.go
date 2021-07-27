package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	vip := viper.New()
	vip.SetDefault("HOST", "0.0.0.0")
	vip.SetDefault("PORT", "8080")
	vip.SetDefault("CONSUL_ADDR", "http://127.0.0.1:8500")
	vip.AutomaticEnv()

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
