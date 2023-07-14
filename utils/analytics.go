package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/todzuko/analytics-service/database/models"
	"github.com/todzuko/analytics-service/workerpool"
	"net/http"
	"time"
)

var Router *chi.Mux

type request struct {
	Module    string            `json:"module"`
	EventType string            `json:"type"`
	Event     string            `json:"event"`
	Name      string            `json:"name"`
	Data      map[string]string `json:"data"`
}

func Construct(r *chi.Mux) {
	Router = r
}

func PostAnalytics(w http.ResponseWriter, r *http.Request) {
	analytics := models.AnalyticsData{
		"",
		r.Header.Get("X-Tantum-Authorization"),
		time.Now(),
		make(map[string]interface{}),
	}
	req := new(request)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return
	}

	analytics.Data["headers"] = r.Header
	analytics.Data["body"] = &req

	workerpool.QueueAnalytics(analytics)
	fmt.Println(analytics)
	fmt.Println(r)
	fmt.Println("response sent")
	respondWithJSON(w, http.StatusAccepted, map[string]string{"status": "OK"})
}

func GetAnalytics(w http.ResponseWriter) {
	analyticsList := models.GetList()
	respondWithJSON(w, http.StatusOK, analyticsList)
}

func GetAnalyticsById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	analyticsValue := models.GetById(id)
	if analyticsValue.ID == "" {
		respondWithJSON(w, http.StatusNotFound, nil)
		return
	}
	respondWithJSON(w, http.StatusOK, analyticsValue)
}
