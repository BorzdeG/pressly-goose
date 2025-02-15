package gomigrations

import (
	"database/sql"

	"github.com/BorzdeG/pressly-goose/v4"
)

func init() {
	goose.AddMigrationNoTx(up001, nil)
}

func up001(db *sql.DB) error {
	q := "CREATE TABLE foo (id INTEGER)"
	_, err := db.Exec(q)
	return err
}
