package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func OpenConnection() (*sql.DB, error) {
	datasource := fmt.Sprintf("user=%s dbname=%s "+
		"password=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
	fmt.Println(datasource)
	return sql.Open("postgres", datasource)
}
