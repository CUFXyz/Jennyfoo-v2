package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBInstance struct {
	sqlinstance *sqlx.DB
}

func Connect() *DBInstance {
	db, err := sqlx.Open("postgres", os.Getenv("LOGIN"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	return &DBInstance{
		sqlinstance: db,
	}
}

func (dbi *DBInstance) GET() {

}
