package config

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(connectionString string) (*sql.DB, error) {
	return sql.Open("mysql", connectionString)
}

func InitDBTest(dbName string, connectionString string) (*sql.DB, error) {
	dbTest, err := sql.Open("mysql", connectionString)

	_, err = dbTest.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = dbTest.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		panic(err)
	}

	query, err := ioutil.ReadFile("../ddl_test.sql")
	if err != nil {
		panic(err)
	}
	if _, err := dbTest.Exec(string(query)); err != nil {
		panic(err)
	}

	_, err = dbTest.Exec(`INSERT INTO ` + dbName + `.category_product (description) VALUES ("Anak");`)
	if err != nil {
		panic(err)
	}

	_, err = dbTest.Exec(`INSERT INTO ` + dbName + `.transaction_status (description) VALUES ("Berhasil");`)
	if err != nil {
		panic(err)
	}

	return dbTest, err
}
