package datastore

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/hyuti/clean-slice-template/services/service-1/ent"
	_ "github.com/mattn/go-sqlite3"
)

func OpenClient(databaseUrl string, maxPoolSize int) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	db.SetMaxOpenConns(maxPoolSize)

	// Create an ent.Driver from `db`.
	return ent.NewClient(ent.Driver(drv)), nil
}

func NewClient(url string, poolMax int) (*ent.Client, error) {
	return OpenClient(url, poolMax)
}
