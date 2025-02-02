package dialectquery

import "fmt"

type Mysql struct{}

var _ Querier = (*Mysql)(nil)

func (m *Mysql) CreateTable(tableName string) string {
	q := `CREATE TABLE %s (
		id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		version_id bigint NOT NULL,
		is_applied boolean NOT NULL,
		tstamp timestamp NULL default now(),
		PRIMARY KEY(id)
	)`
	return fmt.Sprintf(q, tableName)
}

func (m *Mysql) TableExists(_ string) string {
	// TODO https://github.com/pressly/goose/issues/898
	return ""
}

func (m *Mysql) InsertVersion(tableName string) string {
	q := `INSERT INTO %s (version_id, is_applied) VALUES (?, ?)`
	return fmt.Sprintf(q, tableName)
}

func (m *Mysql) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id=?`
	return fmt.Sprintf(q, tableName)
}

func (m *Mysql) GetMigrationByVersion(tableName string) string {
	q := `SELECT tstamp, is_applied FROM %s WHERE version_id=? ORDER BY tstamp DESC LIMIT 1`
	return fmt.Sprintf(q, tableName)
}

func (m *Mysql) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied from %s ORDER BY id DESC`
	return fmt.Sprintf(q, tableName)
}

func (m *Mysql) GetLatestVersion(tableName string) string {
	q := `SELECT MAX(version_id) FROM %s`
	return fmt.Sprintf(q, tableName)
}
