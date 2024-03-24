package main

import (
	"context"
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
	"log"
	"os"
	"time"
)

func main() {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	row := db.QueryRowContext(context.Background(), "SELECT name FROM Mixology.Mixology.Ingredients")
	var name string
	err = row.Scan(&name)
	if err != nil {
		log.Fatalf("could not scan rows of query: %s", err)
	}
	log.Printf("%s", name)
}
