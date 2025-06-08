package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/todo")
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
