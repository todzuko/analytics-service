package models

import (
	"context"
	"fmt"
	"github.com/todzuko/analytics-service/database"
	"log"
	"time"
)

type AnalyticsData struct {
	ID        string                 `json:"id"`
	UserID    string                 `json:"user_id"`
	CreatedAt time.Time              `jsom:"created_at"`
	Data      map[string]interface{} `json:"data"`
}

func GetList() []*AnalyticsData {
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM analytics_table")
	if err != nil {
		log.Fatal("could not execute query ", err)
	}
	defer rows.Close()

	var analyticsList []*AnalyticsData

	for rows.Next() {
		analytics := *new(AnalyticsData)
		err := rows.Scan(&analytics.ID, &analytics.UserID, &analytics.CreatedAt, &analytics.Data)
		if err != nil {
			log.Fatal("could not read from db ", err)
		}

		analyticsList = append(analyticsList, &analytics)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("could not process data ", err)
	}

	return analyticsList
}

func GetById(id string) *AnalyticsData {
	query := fmt.Sprintf("SELECT * FROM analytics_table WHERE id='%s'", id)
	row := database.DB.QueryRow(context.Background(), query)

	analytics := *new(AnalyticsData)
	err := row.Scan(&analytics.ID, &analytics.UserID, &analytics.CreatedAt, &analytics.Data)
	if err != nil {
		log.Println("error response from db ", err)
	}
	return &analytics
}

func Save(analyticsData *AnalyticsData) {
	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		log.Println("could not open db connection")
		return
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO analytics_table (user_id, data, created_at) VALUES ($1, $2, $3)",
		analyticsData.UserID, analyticsData.Data, analyticsData.CreatedAt)
	if err != nil {
		log.Println("Could not execute insert query")
		return
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("Could not commit transaction")
		return
	}

}
