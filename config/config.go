package config

import (
	"database/sql"
	"fmt"
)

func InitDB(connectionString string) (*sql.DB, error) {
	fmt.Println(connectionString)
	return sql.Open("mysql", connectionString)
}
