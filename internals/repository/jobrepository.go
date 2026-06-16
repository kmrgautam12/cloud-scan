package repository

import "github.com/jackc/pgx/v5/pgxpool"

type JobRepository struct {
	db *pgxpool.Pool
}

func NewJobRepository(db *pgxpool.Pool) *JobRepository {
	return &JobRepository{
		db: db,
	}
}
