package dialectquery

import "github.com/BorzdeG/pressly-goose/v4/internal/dialect"

type Turso struct {
	Sqlite3
}

var _ Querier = (*Turso)(nil)

func (t *Turso) GetDialect() dialect.Dialect { return dialect.Turso }
