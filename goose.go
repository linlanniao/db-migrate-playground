package main

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "github.com/pressly/goose/v3"
)

//go:embed examples/migrations2/*.sql
var embedMigrations embed.FS

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "example"
)

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var db *sql.DB
	var err error
	// setup database
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Status(db, "examples/migrations2"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "examples/migrations2"); err != nil {
		panic(err)
	}

	if err := goose.Status(db, "examples/migrations2"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, "examples/migrations2"); err != nil {
		panic(err)
	}
	if err := goose.Status(db, "examples/migrations2"); err != nil {
		panic(err)
	}

	if err := goose.Reset(db, "examples/migrations2"); err != nil {
		panic(err)
	}

	// run app
}
