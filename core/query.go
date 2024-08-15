package core

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type ConnectionConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func ProcessQuery(configs []ConnectionConfig, query string) {}

func ExecuteQuery(config ConnectionConfig, query string) ([]interface{}, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", config.Username, config.Database, config.Password, config.Host, strconv.Itoa(config.Port))

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var result []interface{}

	for rows.Next() {
		var now string
		err = rows.Scan(&now)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, now)
	}

	return result, nil
}
