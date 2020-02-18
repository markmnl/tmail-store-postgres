package pgstore

import (
	"fmt"
	"log"
	"github.com/markmnl/tmail-store/tstore/pkg"
	"database/sql"
	"github.com/lib/pq"
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

	connStr := "" // https://www.postgresql.org/docs/current/libpq-envars.html
	defaultDb, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return defaultDb
}

func toNullString(s *string) sql.NullString {
    if len(*s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
         String: *s,
         Valid: true,
    }
}

// ParentExists queries database for existance of message with sha256 supplied
func ParentExists(pSHA256 []byte) (bool, error) {
	db := acquireConn()
	exists := false

	dbErr := db.QueryRow("select exists(select 1 from msg where sha256 = $1)", pSHA256).Scan(&exists)
	if dbErr != nil {
		return false, fmt.Errorf("Failed to query database: %w", dbErr)
	}

	return exists, nil
}

// Store stores message supplied in postgres db
func Store(msg *tstore.Msg) error {

	db := acquireConn()

	id := msg.ID[:]
	var pid []byte = nil
	if msg.PID64 != "" {
		pid = msg.PID[:]
	}

	_, dbErr := db.Exec(`insert into msg(ts
			, from_addr
			, to_addr
			, topic
			, type
			, sha256
			, psha256
			, content)
		values ($1, $2, $3, $4, $5, $6, $7, $8)`, msg.Time,
			msg.From, 
			(*pq.StringArray)(&msg.To),
			msg.Topic,
			msg.Type,
			id, 
			pid,
			msg.Content)
	if dbErr != nil {
		return fmt.Errorf("Database error: %w", dbErr)
	}

	// TODO attachments

	return nil
}