package internal

import "github.com/pressly/goose/v3"

func GenerateMigration(name string, migrationType string) error {
	if migrationType == "" {
		migrationType = "sql"
	}

	return goose.Create(nil, "migrations", name, migrationType)
}
