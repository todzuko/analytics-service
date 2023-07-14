package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type DbConnection struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var DB *pgx.Conn

func Connect() {
	dbInfo := DbConnection{
		"postgres",
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	}

	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbInfo.user, dbInfo.password, dbInfo.host, dbInfo.port, dbInfo.dbname,
	)
	var err error
	DB, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = CreateTableIfNotExists()
	if err != nil {
		return
	}
	log.Println("Connected to database")
}

func CreateTableIfNotExists() error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS analytics_table (
			id SERIAL PRIMARY KEY,
			user_id VARCHAR(64),
			created_at TIMESTAMP,
			data JSONB
		)
	`

	_, err := DB.Exec(context.Background(), createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table %v", err)
	}

	log.Println("created table")
	return nil
}
