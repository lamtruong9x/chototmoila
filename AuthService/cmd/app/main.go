package main

import (
	controller "AuthService/internal/controller"
	repository "AuthService/internal/db/sqlc"
	"AuthService/internal/token"
	"database/sql"
	"flag"
	"fmt"
	log2 "log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr string
	dsn  string
}

type application struct {
	controller controller.IController
}

var cfg config

const secretKey = "3y6B8DbGdJfNjQmSqVsXu2x4z7C9EbHeKgNr"

func main() {
	flag.StringVar(&cfg.addr, "addr", ":4040", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "root:pass@tcp(localhost:3030)/cho_tot?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Create db instance
	db, err := OpenDB(cfg.dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// Create Queries instance
	queries := repository.New(db)

	// Create a logger
	logger := log2.New(os.Stderr, "ERROR", log2.Ldate|log2.Ltime)

	// Create token maker instance
	maker, err := token.NewJWTMaker(secretKey)
	//payload, err := maker.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImRhNDM1ODAxLTkyYWMtNDg4OS1iZmE3LTM0MWYyZGEyZjJmNiIsInBob25lIjoiMDM4Nzg1NTUwMCIsInVzZXJfaWQiOjEsImlzQWRtaW4iOmZhbHNlLCJpc3N1ZWRfYXQiOiIyMDIyLTA3LTA3VDEyOjQ2OjU3LjM5Nzk2NjgrMDc6MDAiLCJleHBpcmVkX2F0IjoiMjAyMi0wNy0wN1QxMzozMTo1Ny4zOTc5NjY4KzA3OjAwIn0.qaiC7TKmVpShvwnr7G6hhE6xuY49wZH82XgPDK0Y_zM")
	logger.Println(err)

	// Create controller instance
	controller := controller.New(queries, maker, logger)

	// Create app
	app := &application{controller: controller}
	srv := app.NewRouter()
	fmt.Println("Starting app on port:", cfg.addr)
	panic(srv.Run(cfg.addr))
}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
