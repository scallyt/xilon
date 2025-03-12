package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var (
	Ctx  = context.Background()
	Conn *pgx.Conn
)

func ConnectDB() error {
	dns := "postgres://neko:neko123@localhost:5432/nekodb?sslmode=disable"

	var err error
	Conn, err = pgx.Connect(Ctx, dns)
	if err != nil {
		return err
	}
	return nil
}
