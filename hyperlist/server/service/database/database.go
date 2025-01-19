package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type (
	DBConfig struct {
		// should only be used for select statements
		*sql.DB

		// should only be used for write statements
		write *sql.Conn
	}

	WriteOptions struct {
		Query   string
		Args    []any
		Timeout time.Duration
	}
)

func (cfg DBConfig) Close() {

	// optimise for future queries
	_, err := cfg.DB.Exec("PRAGMA optimize")
	if err != nil {
		log.Print(err)
	}

	if err := cfg.write.Close(); err != nil {
		log.Print(err)
	}
	if err := cfg.DB.Close(); err != nil {
		log.Print(err)
	}
}

func (cfg DBConfig) WriteExec(opts WriteOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout*time.Second)
	_, err := cfg.write.ExecContext(ctx, opts.Query, opts.Args...)
	cancel()
	return err
}

func (cfg DBConfig) WriteQuery(opts WriteOptions) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout*time.Second)
	rows, err := cfg.write.QueryContext(ctx, opts.Query, opts.Args...)
	cancel()
	return rows, err
}

func CreateDBConfig() DBConfig {
	db, err := sql.Open("sqlite3", "./database/hyperlist.db")
	if err != nil {
		panic(err)
	}

	// ping database, to check is exists
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// write to a write-ahead-log and regularily commit the changes, rather than db file
	// allows multiple concurrent readers
	if _, err = db.Exec("pragma journal_mode = WAL;"); err != nil {
		panic(err)
	}
	// recommened option when using WAL
	if _, err = db.Exec("pragma synchronous = normal;"); err != nil {
		panic(err)
	}

	// checks if table exists
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='sites';")
	if err != nil {
		panic(err)
	}

	// we expect a single string since there can only be one
	var site string
	if rows.Next() {
		if err = rows.Scan(&site); err != nil {
			panic(err)
		}
	}

	// if table doesn't exist, create the table
	if site != "sites" {
		sqlStmt := `create table sites (id integer not null primary key, 
										slug text, 
										name text, 
										src text, 
										url text, 
										created integer, 
										alive tinyint);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			panic(err)
		}
	}

	// create writer connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}
	cancel()

	return DBConfig{
		DB:    db,
		write: conn,
	}
}
