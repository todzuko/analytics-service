package main

import (
	"fmt"
	"github.com/todzuko/analytics-service/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Router *chi.Mux

func main() {
	Router = chi.NewRouter()
	Router.Use(middleware.Logger)
	database.Connect()

	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("Port not given")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), Router))
}
