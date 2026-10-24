package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
  return pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres")
}
