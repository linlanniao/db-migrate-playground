package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://examples/migrations1",
		"postgres://postgres:postgres@localhost:5432/example?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Version())

	if err := m.Up(); err != nil {
		log.Println(err)
	}
	fmt.Println(m.Version())

	if err := m.Down(); err != nil {
		log.Println(err)
	}
	fmt.Println(m.Version())
}
