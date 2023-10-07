package database

import (
	"anote/internal/types"
	errors "anote/internal/types"
	"database/sql"
	"log"
	"reflect"
	"sync"

	"github.com/lib/pq"
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

		log.Println("Connected to PostgreSQL!")
		conn = Conn{conn: c}
	})
	return conn

}

func getFields(element reflect.Value) []interface{} {
	fields := make([]interface{}, element.Elem().NumField())
	for i := 0; i < element.Elem().NumField(); i++ {
		fields[i] = element.Elem().Field(i).Addr().Interface()
	}
	return fields
}

// Returns only the first result of the query
func (c Conn) Exec(query string, args ...any) *errors.AppError {
	_, err := c.conn.Exec(query, args...)
	if err != nil {
		log.Println("[DBConn] Exec error: ", err)
		if e := err.(*pq.Error); e.Code == "23505" {
			return types.NewAppError(400, "Resource already exists")
		}
		return types.NewAppError(500, "Internal server error")
	}
	return nil
}

func (c Conn) QueryOne(objType reflect.Type, query string, args ...any) (any, *errors.AppError) {
	queryResult := c.conn.QueryRow(query, args...)
	if queryResult.Err() != nil {
		log.Println("[DBConn] QueryOne query error: ", queryResult.Err())
		return nil, errors.NewAppError(500, "Internal server error")
	}

	element := reflect.New(objType)
	if err := queryResult.Scan(getFields(element)...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("[DBConn] QueryOne scan error: ", err)
		return nil, errors.NewAppError(500, "Internal server error")
	}
	return element.Elem().Interface(), nil
}

func (c Conn) QueryMultiple(objType reflect.Type, query string, args ...any) (any, *errors.AppError) {
	queryResult, err := c.conn.Query(query, args...)
	if err != nil {
		log.Println("[DBConn] QueryMultiple query error: ", err)
		return nil, errors.NewAppError(500, "Internal server error")
	}

	sliceType := reflect.SliceOf(objType)
	elements := reflect.MakeSlice(sliceType, 0, 0)
	defer queryResult.Close()
	for queryResult.Next() {
		element := reflect.New(objType)
		if err := queryResult.Scan(getFields(element)...); err != nil {
			log.Println("[DBConn] QueryMultiple scan error: ", err)
			return nil, errors.NewAppError(500, "Internal server error")
		}
		elements = reflect.Append(elements, element.Elem())
	}
	return elements.Interface(), nil
}
