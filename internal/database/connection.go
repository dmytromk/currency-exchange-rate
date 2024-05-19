package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func OpenConnection() (*sql.DB, error) {
	datasource := fmt.Sprintf("host=%s user=%s dbname=%s "+
		"password=%s port=5432 sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
	return sql.Open("postgres", datasource)
}
