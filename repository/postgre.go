package repository

import "database/sql"

type PostgreRepository struct {
	DB *sql.DB
}

func NewPostgreRepository(conn string) *PostgreRepository {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	return &PostgreRepository{DB: db}
}
