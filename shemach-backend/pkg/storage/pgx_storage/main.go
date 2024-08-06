package pgx_storage

import (
	"context"
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"

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
	tm.Println(tm.Color(" pgx : DB Connected Succesfuly ... \n", tm.GREEN))

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				{
					if conn != nil {
						err = conn.Ping(context.Background())
					}
					if err != nil || conn == nil {
						tm.Println(tm.Color("DB Connection: Trying to reconnect ...", tm.RED))
						i := 0
						for ; i <= 100; i++ {
							conn, err = pgxpool.Connect(context.Background(), postgresStatment)
							if err == nil {
								break
							}
						}
						if i == 100 {
							tm.Println(tm.Color("Database connection Failure ...", tm.RED))
							os.Exit(1)
						}
					}
				}
			}
		}
	}()

	return conn, err
}
