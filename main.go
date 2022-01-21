package main

import (
	"fmt"
	"os"
	"sirclo/api/config"
	"sirclo/api/router"
)

func main() {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	fmt.Println(connectionString)
	db, err := config.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	e := router.InitRoute(db)
	e.Logger.Fatal(e.Start(":8080"))

}
