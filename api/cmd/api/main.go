package main

import (
	"ft"
	"net/http"

	"githhub.com/go-chi/chi"
	"github.com/mauriciomolina/lets-go/api/internal/handlers"
	log "github.com/sirupsen/logrus"
)


func main () {
	log.SetReporCaller(true)
	var r *chi.Mux = chi.NewRouter() //pointer to mux type -> struct para setear api
	handlers.Hanlder(r)

	fmt.Println("STARTING API")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("error starting server:", err)
	}
}