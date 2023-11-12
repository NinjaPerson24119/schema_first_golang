package database

import (
	"context"
	"fmt"

	"github.com/NinjaPerson24119/schema_first_price_history/internal/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDatabase interface {
	GetConnection(ctx context.Context) (*pgxpool.Conn, error)
	GetQueries(ctx context.Context) (*sqlc.Queries, error)
	Close()
}

type Database struct {
	// thread-safe
	pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connectionURL string) (IDatabase, error) {
	config, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres connection url: %v", err)
	}
	config.MinConns = 5
	config.MaxConns = 25
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %v", err)
	}

	return &Database{
		pool: pool,
	}, nil
}

func (s *Database) GetConnection(ctx context.Context) (*pgxpool.Conn, error) {
	return s.pool.Acquire(ctx)
}

func (s *Database) GetQueries(ctx context.Context) (*sqlc.Queries, error) {
	conn, err := s.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	return sqlc.New(conn), nil
}

func (s *Database) Close() {
	s.pool.Close()
}
