package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type DBInstance struct {
	Sqlinstance *sqlx.DB
}

func Connect() *DBInstance {
	db, err := sqlx.Open("postgres", os.Getenv("LOGIN"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	return &DBInstance{
		Sqlinstance: db,
	}
}
