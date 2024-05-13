package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // Use pgx as the driver
)

var DB *sql.DB // Corrected variable name

func SetupDatabase() {
	var err error
	DB, err = sql.Open("pgx", Conf.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Configure the connection pool
	DB.SetMaxIdleConns(Conf.Database.MaxIdleConns)
	DB.SetMaxOpenConns(Conf.Database.MaxOpenConns)
	connMaxLifetime, err := time.ParseDuration(Conf.Database.ConnMaxLifetime)
	if err != nil {
		log.Fatalf("Error parsing duration (ConnMaxLifetime): %v", err)
	}
	DB.SetConnMaxLifetime(connMaxLifetime)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
}
