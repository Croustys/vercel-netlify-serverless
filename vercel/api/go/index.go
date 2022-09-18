package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectToDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "POST":
		if create(r) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Successful creation"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unsuccesful database write"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Unsupported Method"))
	}
}

type reqBody struct {
	Name string `json:"name"`
}

func create(r *http.Request) bool {
	db := connectToDb()
	defer db.Close()

	var rq reqBody
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		log.Println(err)
	}

	result, err := db.Exec("INSERT INTO test_db (name) VALUES ($1)", rq.Name)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)

	return rows == 1
}
