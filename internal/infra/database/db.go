package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func New() *pgx.Conn {
	url := os.Getenv("DATABASE_URL")
	fmt.Printf("url: %v\n", url)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	return conn
}
