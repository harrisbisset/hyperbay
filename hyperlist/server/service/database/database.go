package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	*sql.DB
}

func CreateDBConfig() DBConfig {
	db, err := sql.Open("sqlite3", "./hyperlist.db")
	if err != nil {
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

	return DBConfig{
		DB: db,
	}
}
