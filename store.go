package burrow

import (
	"context"
	"encoding/json"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"

	"github.com/jackc/pgx/v5/pgxpool"
)

const name = "github.com/jenmud/burrow"

var (
	tracer = otel.Tracer(name)
	meter  = otel.Meter(name)
	logger = otelslog.NewLogger(name)
)

const (
	connectTimeout = 30 * time.Second
	closeTimeout   = 30 * time.Second
)

// Store simplified layer used for querying the underlying data.
type Store struct {
	db *pgxpool.Pool
}

// NewStore return a new store.
func NewStore(ctx context.Context, dsn string) (*Store, error) {
	ctx, cancel := context.WithTimeout(ctx, connectTimeout)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Store{db: pool}, nil
}

// Close closes the store.
func (s *Store) Close() error {
	s.db.Close()
	return nil
}

// DefaultLimit is the default limit for queries.
const DefaultLimit = 100

type NodesArgs struct {
	Limit  uint
	Cursor string
}

// Nodes returns all the nodes in the store.
func (s *Store) Nodes(ctx context.Context, args NodesArgs) ([]Node, error) {
	ctx, span := tracer.Start(ctx, "Nodes")
	defer span.End()

	if args.Limit == 0 {
		args.Limit = DefaultLimit
	}

	query := `
		SELECT
			n.id,
			n.properties,
			COALESCE(array_agg(l.name ORDER BY l.name) FILTER (WHERE l.name IS NOT NULL), '{}') AS labels
		FROM nodes n
		LEFT JOIN node_labels l ON l.node_id = n.id
		WHERE n.id > $1
		GROUP BY n.id
		ORDER BY n.id ASC
		LIMIT $2;
	`

	rows, err := s.db.Query(ctx, query, args.Cursor, args.Limit)
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0, args.Limit)

	for rows.Next() {
		n := Node{}

		labels := []byte{}
		props := []byte{}

		if err := rows.Scan(&n.ID, &props, &labels); err != nil {
			return nodes, err
		}

		if err := json.Unmarshal(props, &n.Properties); err != nil {
			return nodes, err
		}

		if err := json.Unmarshal(labels, &n.Labels); err != nil {
			return nodes, err
		}
	}

	return nodes, nil
}
