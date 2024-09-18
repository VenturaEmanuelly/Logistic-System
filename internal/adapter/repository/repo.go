package repository

import "database/sql"

type RepositSql struct {
	db *sql.DB
}

func (r RepositSql) QueryRow(query string , args []any, dest ...any) error {
	return r.db.QueryRow(query, args ...).Scan(dest...)
}


func NewRepoSql (db *sql.DB) RepositSql {
	return RepositSql{db: db}
}