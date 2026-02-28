package localdb

import (
	"log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"git.dilangilluly.us/dbob16/tam4/apiserver/settings"
)

func NewSession() (*sql.DB) {
	config := settings.ReadSettings()
	connstr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBCreds.DBUser, config.DBCreds.DBPassword, config.DBCreds.DBHost, config.DBCreds.DBPort, config.DBCreds.DBDatabase)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Println(err)
	}
	return db
}

func InitDB() (string) {
	db := NewSession()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatalln("Unable to initiate database due to a connection error: ", err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS api_keys (api_key VARCHAR(255) PRIMARY KEY, description VARCHAR(255))")
	db.Exec("CREATE TABLE IF NOT EXISTS prefixes (prefix VARCHAR(255) PRIMARY KEY, color VARCHAR(100), weight INT)")

	return "DB Initialized Successfully"
}
