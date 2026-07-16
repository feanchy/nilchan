package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgres://postgres:1991@localhost:5432/postgres")
}
