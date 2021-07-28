package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
	if vip.GetString("ADDR") != "" {
		addr = vip.GetString("ADDR")
	} else {
		addr = fmt.Sprintf("%s:%s", vip.GetString("HOST"), vip.GetString("PORT"))
	}

	log.Println("Server listen at:", addr)
	log.Fatalln(http.ListenAndServe(addr, r))
}
