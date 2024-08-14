package core

import "log"

type ConnectionConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func ProcessQuery(configs []ConnectionConfig, query string) {

}

func ExecuteQuery(config ConnectionConfig, query string) (interface{}, error) {
	log.Printf("Executing query: %s on database: %s:%v with user: %s and password: %s", query, config.Host, config.Port, config.Username, config.Password)
	return nil, nil
}
