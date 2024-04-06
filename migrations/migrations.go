package migrations

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(dsn string) {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(b), "migrations")

	m, err := migrate.New("file://"+dir, dsn)
	if err != nil {
		log.Fatalf("Error while creating migrate instance: %v", err)
	}

	// Apply all the available migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error while applying migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
