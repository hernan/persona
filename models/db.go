package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var MyDb *sql.DB

func DBInit() {
	db, err := sql.Open("mysql", "hernan:changeme@/persona_development")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	MyDb = db
}
