package main

import (
	"database/sql"
	"log"

	"git.01.alem.school/Kusbek/social-network/backend/pkg/db/sqlite"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func main() {
	sqlOpts := &sqlite.Options{
		Address: "local.db",
	}

	db, err := sqlite.Init(sqlOpts)
	if err != nil {
		log.Fatalf("Failed to start db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()

	initMigrate(db)
}

func initMigrate(db *sql.DB) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://pkg/db/migrations/sqlite", "ql", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
