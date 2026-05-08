package burrowdb

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	connectTimeout = 30 * time.Second
	closeTimeout   = 30 * time.Second
)

// Store simplified layer used for querying the underlying data.
type Store struct {
	db *pgx.Conn
}

// NewStore return a new store.
func NewStore(ctx context.Context, dsn string) (*Store, error) {
	ctx, cancel := context.WithTimeout(ctx, connectTimeout)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Store{db: conn}, nil
}

// Close closes the store.
func (s *Store) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), closeTimeout)
	defer cancel()
	return s.db.Close(ctx)
}
