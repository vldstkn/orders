package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sqlx.DB
}

func NewDb(connectString string) *DB {
	fmt.Println(connectString)
	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	return &DB{
		db,
	}
}
