package database

import (
	"context"
	"distributed-job-scheduler/jobqueue/utils/helper"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	host     = helper.GetStringFromEnv("HOST")
	port     = helper.GetStringFromEnv("PORT")
	user     = helper.GetStringFromEnv("USER")
	password = helper.GetStringFromEnv("PASSWORD")
	dbName   = helper.GetStringFromEnv("DB_NAME")
)
var connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)

func InitialiseSchedulerStore() (pgx *pgxpool.Pool, err error) {
	ctx := context.Background()
	pgx, err = newPostgresConnection(ctx, connectionString)
	return pgx, err
}

func newPostgresConnection(e context.Context, connectionStr string) (*pgxpool.Pool, error) {
	return pgxpool.New(e, connectionStr)
}
