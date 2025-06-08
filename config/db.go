package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	host := os.Getenv("PG_HOST")
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://postgres:postgres@%v:5432/todo?sslmode=disable", host))
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}
	DB = conn

	_, err = DB.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE
		)
	`)
	if err != nil {
		panic(fmt.Sprintf("Unable to create table: %v", err))
	}
	fmt.Println("Connected to PostgreSQL")
}
