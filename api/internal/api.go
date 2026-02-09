package handlers

import (
	"github.com/go-chi/chi"
	chimiddle 	
	"github.com/mauriciomolinapicco/lets-go/api/internal/middlewre"
)


func Handler(r *chi.Mux) {
	//global middleware
	r.Use(chimiddle.StripSlashes) //elimina las barras al final de las rutas
	r.Route("/account", func(router chi.Router)) {

		router.Use(middleware.Authorization)
		
		router.Get("/coins", GetCoinBalance)

	}

}