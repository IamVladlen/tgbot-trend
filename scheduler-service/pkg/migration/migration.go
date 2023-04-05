package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// LoadMigrationSQL applies up migration from .sql file
// to postgres database.
func LoadMigrationSQL(dbURL string, sourceURL string) {
	migration, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		log.Fatalln("Cannot run database migration:", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("Cannot run database migration:", err)
	}
}
