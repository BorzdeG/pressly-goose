package register

import (
	"database/sql"

	"github.com/BorzdeG/pressly-goose/v4"
)

func init() {
	goose.AddMigration(
		func(_ *sql.Tx) error { return nil },
		func(_ *sql.Tx) error { return nil },
	)
}
