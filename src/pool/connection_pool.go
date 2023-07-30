package pool

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func CreatePool() {
	p, err := pgxpool.New(context.Background(), "postgresql://postgres:123123123@localhost:5432/library?sslmode=disable&pool_max_conns=10")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	pool = p
}

func DestroyPool() {
	pool.Close()
}

func GetConnection() (*pgxpool.Conn, error) {
	return pool.Acquire(context.Background())
}
