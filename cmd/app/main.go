package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mmunareto/zipcode-api/internal/infra/webserver/handlers"
	"github.com/mmunareto/zipcode-api/internal/services"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	zipCodeService := services.NewZipCodeService(http.DefaultClient)
	zipCodeHandler := handlers.NewZipCodeHandler(zipCodeService)

	r.Get("/zip-details/{zipCode}", zipCodeHandler.GetZipCodeDetails)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}