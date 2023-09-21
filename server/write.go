package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Define a PostgreSQL connection string
	connStr := "user=anote password=anote dbname=anote sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")

	// Now you can use the 'db' variable to execute SQL queries.
	_, err = db.Exec("INSERT INTO users (id, email, password) VALUES ('danihtoledo', 'email@email.com', 'senhaSegura')")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO notes (id, title, author_id, content) VALUES ('550e8400-e29b-41d4-a716-446655440000', 'Note Title', 'danihtoledo', 'Note Content')")
	if err != nil {
		log.Fatal(err)
	}
}
