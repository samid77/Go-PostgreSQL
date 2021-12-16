package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-crud/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
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

// Create User in Postgres database
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//create empty user of type models.user
	var user models.User

	//decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	//call insert user function and pass the user
	insertID := insertUser(user)

	//format a response object
	res := response{
		ID:      insertID,
		Message: "User Created Successfully",
	}

	//send the response
	json.NewEncoder(w).Encode(res)
}

func insertUser(user models.User) int64 {

	db := createConnection()

	defer db.Close()

	//sql query for insert data
	//returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`
	var id int64

	//execute sql statement
	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute sql statement: %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}
