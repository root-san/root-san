package impl

import (
	_ "github.com/go-sql-driver/mysql" // import driver
	"github.com/jmoiron/sqlx"
	"github.com/root-san/root-san/app/repository"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository.Repository {
	return &Repository{db: db}
}
