package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var dbURI, migrationsPath, migrationsTable, direction string

	flag.StringVar(&dbURI, "db-uri", "", "connection string to db")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&direction, "direction", "up", "direction of migrations: up or down")
	// optional
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if dbURI == "" {
		panic("storage-path is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		dbURI,
	)
	if err != nil {
		panic(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")

				return
			}

			panic(err)
		}

		fmt.Println("migrations applied successfully")

	case "down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")

				return
			}

			panic(err)
		}

		fmt.Println("migrations rolled back successfully")

	default:
		log.Fatalf("invalid direction: %s", direction)
	}

	fmt.Println("migrations applied successfully")
}
