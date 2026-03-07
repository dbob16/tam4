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
	db.Exec(`CREATE TABLE IF NOT EXISTS tickets (
		prefix VARCHAR(255),
		ticket_id INT,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		phone_number VARCHAR(255),
		preference VARCHAR(20),
		PRIMARY KEY (prefix, ticket_id)
		)`)
	db.Exec(`CREATE TABLE IF NOT EXISTS baskets (
		prefix VARCHAR(255),
		basket_id INT,
		description VARCHAR(255),
		donors VARCHAR(255),
		winning_ticket INT NOT NULL DEFAULT 0,
		PRIMARY KEY (prefix, basket_id)
		)`)
	db.Exec(`CREATE OR REPLACE VIEW drawing AS
		SELECT b.prefix, b.basket_id, b.description, b.winning_ticket, CONCAT(t.last_name, ', ', t.first_name) AS winner_name,
		t.phone_number FROM baskets b LEFT JOIN tickets t ON b.prefix = t.prefix AND b.winning_ticket = t.ticket_id
		ORDER BY b.prefix ASC, b.basket_id`)

	return "DB Initialized Successfully"
}
