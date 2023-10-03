package database

import (
	interfaces "anote/internal/interfaces/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

// This component is carried with connecting to the database and sending the queries
type Conn struct {
	conn *sql.DB
}

var conn Conn
var once sync.Once

func GetConnection() Conn {
	once.Do(func() {
		connStr := "user=anote password=anote dbname=anote sslmode=disable"
		c, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		if err = c.Ping(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to PostgreSQL!")
		conn = Conn{conn: c}
	})
	return conn
}

// Returns only the first result of the query
func (c Conn) QueryOne(dest interfaces.Entity, query string, args ...any) error {
	queryResult := c.conn.QueryRow(query, args...)

	if err := queryResult.Scan(dest.GetFieldAdresses()...); err != nil {
		errMessage := fmt.Sprintf("Error parsing single query result: %v", err)
		log.Println(errMessage)
		return errors.New(errMessage)
	}

	return nil
}

func (c Conn) QueryMultiple(dest []interfaces.Entity, query string, args ...any) error {
	queryResult, err := c.conn.Query(query, args...)
	if err != nil {
		errMessage := fmt.Sprintf("Error on query: %v", err)
		log.Println(errMessage)
		return errors.New(errMessage)
	}

	defer queryResult.Close()
	for queryResult.Next() {
		if err := queryResult.Scan(dest[0].GetFieldAdresses()...); err != nil {
			errMessage := fmt.Sprintf("Error parsing multiple query result: %v", err)
			log.Println(errMessage)
			return errors.New(errMessage)
		}
	}
	return nil
}

func (c Conn) Exec(query string, args ...any) error {
	_, err := c.conn.Exec(query, args...)
	if err != nil {
		errMessage := fmt.Sprintf("Error on query exec: %v", err)
		log.Println(errMessage)
		return errors.New(errMessage)
	}
	return nil
}
