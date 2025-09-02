package infrastructure

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

// NewPostgresDB creates a new PostgreSQL connection pool.
func NewPostgresDB(dsn string) *sql.DB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open databse: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database :%v", err)
	}

	return db
}
