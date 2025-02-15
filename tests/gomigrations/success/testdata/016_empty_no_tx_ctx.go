package gomigrations

import (
	"github.com/BorzdeG/pressly-goose/v4"
)

func init() {
	goose.AddMigrationNoTxContext(nil, nil)
}
