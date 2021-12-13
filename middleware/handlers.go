package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//response format
type response struct {
	ID      int64 `json:"id,omitempty"`
	Message int64 `json:"message,omitempty"`
}

//create connection with postgres db
func createConnection() *sql.DB {
	err := godotenv.Load("env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Open the connection
	db, err := sql.Open("postgres", os.Getenv("postgres://dimaseptyanto:  ds@localhost:5432/golangcrud"))

	if err != nil {
		panic(err)
	}

	//check the db connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established")

	return db
}
