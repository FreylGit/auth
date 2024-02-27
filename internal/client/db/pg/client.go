package pg

import (
	"context"
	"github.com/FreylGit/auth/internal/client/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	pg := NewDB(dbc)

	return &pgClient{
		masterDBC: pg,
	}, nil
}

func (p pgClient) DB() db.DB {
	return p.masterDBC
}

func (p pgClient) Close() error {
	if p.masterDBC != nil {
		p.masterDBC.Close()
	}

	return nil
}