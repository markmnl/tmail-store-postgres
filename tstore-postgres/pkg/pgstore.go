package pgstore

import (
	"encoding/hex"
	"fmt"
	"log"
	"github.com/markmnl/tmail-store/tstore/pkg"
	"database/sql"
	_ "github.com/lib/pq" // somehow used by database/sql..
	"github.com/joho/godotenv"
)

var defaultDb *sql.DB = nil


func acquireConn() *sql.DB {
	if defaultDb != nil {
		return defaultDb
	}

	// load env..
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}	

	connStr := "" // defaults loaded from env
	defaultDb, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return defaultDb
}

// Store stores message supplied in postgres db
func Store(msg *tstore.Msg) error {

	db := acquireConn()

	sha456Hex := hex.EncodeToString(msg.ID[:])
	psha256Hex := ""
	if msg.PID64 != "" { // TODO better check..
		psha256Hex = hex.EncodeToString(msg.PID[:])
	}

	_, dbErr := db.Exec(`insert into msg(ts
			, from_addr
			, to_addr
			, topic
			, type
			, sha256
			, psha256
			, content)
		values ($1, $2, $3, $4, $5, $6, $7, $8)`, msg.Time, msg.From, msg.To, msg.Topic,
		msg.Type, sha456Hex, psha256Hex, msg.Content)
	if dbErr != nil {
		return fmt.Errorf("Failed to store msg: %w", dbErr)
	}

	// TODO attachments

	return nil
}