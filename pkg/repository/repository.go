package repository

type Restaurant interface {
}

type Table interface {
}

type Repository struct {
	Restaurant
	Table
}

func NewRepository() *Repository {
	return &Repository{}
}
