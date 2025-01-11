package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AvijitChakraborty1/equity-insights/authentication-service/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication service...")

	log.Println("Start trying to connect to Postgres...")
	con := connectToDB()
	if con == nil {
		log.Fatal("Error connecting to Postgres...")
	}

	app := Config{
		DB:     con,
		Models: data.New(con),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic("Unable to start Authentication service ", err)
	}
}

func openDB() (*sql.DB, error) {
	dbSourceName := os.Getenv("DSN")
	log.Println("dbSourceName : ", dbSourceName)

	db, err := sql.Open("pgx", dbSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	count := 1
	for {
		connection, err := openDB()

		if err != nil {
			log.Println("Still trying to connect to the DB...")
			count++
		} else {
			log.Println("Connected to DB...")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
