package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := initSchema(db); err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

func initSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS agents (
		id TEXT PRIMARY KEY,
		os TEXT,
		arch TEXT,
		last_seen DATETIME
	);
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		agent_id TEXT,
		command TEXT,
		args TEXT,
		status TEXT,
		result TEXT,
		created_at DATETIME,
		FOREIGN KEY(agent_id) REFERENCES agents(id)
	);
	`
	_, err := db.Exec(schema)
	return err
}
