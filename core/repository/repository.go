package repository

import (
	"database/sql"

	ca_parksdb "ca_parks/core/repository/gen"
)

func New(db sql.DB) *ca_parksdb.Queries {
	return ca_parksdb.New(&db)
}
