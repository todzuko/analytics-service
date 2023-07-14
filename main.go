package main

import (
	"fmt"
	"github.com/todzuko/analytics-service/database"
	"github.com/todzuko/analytics-service/utils"
	"github.com/todzuko/analytics-service/workerpool"
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
	utils.Construct(Router)
	workerpool.StartWorkerPool()
	database.Connect()

	Router.Get("/analitycs", func(w http.ResponseWriter, r *http.Request) {
		utils.GetAnalytics(w)
	})
	Router.Post("/analitycs", func(w http.ResponseWriter, r *http.Request) {
		utils.PostAnalytics(w, r)
	})
	Router.Get("/analitycs/{id}", func(w http.ResponseWriter, r *http.Request) {
		utils.GetAnalyticsById(w, r)
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("Port not given")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), Router))
}
