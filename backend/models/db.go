package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Datastore interface {
	AllPostsShorten() ([]Post, error)
}

type DB struct {
	*sql.DB
}

func InitDB(dbUrl string) (*DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return &DB{db}, nil
}
