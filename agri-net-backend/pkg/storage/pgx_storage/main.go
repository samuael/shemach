package pgx_storage

import (
	"context"
	"fmt"
	"os"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewStorage(username, password, host, dbname string) (*pgxpool.Pool, error) {
	// Preparing the statement
	// postgresStatment := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, dbname)
	postgresStatment := os.Getenv("POSTGRES_REMOTE_URI")
	conn, err := pgxpool.Connect(context.Background(), postgresStatment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	print(" pgx : DB Connected Succesfuly ... \n")

	return conn, err
}
