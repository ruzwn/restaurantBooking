package repository

import "github.com/jmoiron/sqlx"

type Restaurant interface {
}

type Table interface {
}

type Repository struct {
	Restaurant
	Table
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
