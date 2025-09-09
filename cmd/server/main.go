package main

import (
	"database/sql"
	"log"
	"net/http"
	"rest-api-postgres/internal/modules/users"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	users.InitDB(db)

	http.HandleFunc("GET /users", users.GetUsers)
	http.HandleFunc("POST /users", users.CreateUser)
	http.HandleFunc("GET /users/{id}", users.GetUserById)
	http.HandleFunc("PUT /users/{id}", users.UpdateUser)

	log.Println("Server starting on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
