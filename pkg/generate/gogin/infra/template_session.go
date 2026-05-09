package infra

// sessionWrapperTemplate is the printf-style template for
// arts/backend/internal/infra/session/postgres.go.
var sessionWrapperTemplate = sessionWrapperHeaderType + sessionWrapperHeaderWhat + `

package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/park-jun-woo/ssac/pkg/session"

	"%[1]s/internal/db"
)

// postgresSession adapts the user's sqlc Queries onto session.SessionModel.
// Construct via NewPostgres(queries); wired from main.go via
// session.Init(infrasession.NewPostgres(queries)).
type postgresSession struct {
	q *db.Queries
}

// NewPostgres returns a session.SessionModel backed by the user's sqlc
// Queries. Keyed state (user_id, cart, etc.) lives in fullend_sessions.
func NewPostgres(q *db.Queries) session.SessionModel {
	return &postgresSession{q: q}
}

// Set marshals value to JSON and upserts with expires_at = now+ttl.
// interface.yaml port: %[2]s.
func (s *postgresSession) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.q.%[2]s(ctx, db.%[2]sParams{
		Key:       key,
		Value:     payload,
		ExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(ttl), Valid: true},
	})
}

// Get forwards to the sqlc query and returns the raw bytes as a string.
// interface.yaml port: %[3]s.
func (s *postgresSession) Get(ctx context.Context, key string) (string, error) {
	raw, err := s.q.%[3]s(ctx, key)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return string(raw), nil
}

// Delete forwards to the sqlc query. interface.yaml port: %[4]s.
func (s *postgresSession) Delete(ctx context.Context, key string) error {
	return s.q.%[4]s(ctx, key)
}
`

var sessionWrapperHeaderType = "//" + "ff:type feature=infra type=model topic=session\n"
var sessionWrapperHeaderWhat = "//" + "ff:what postgresSession — ssac/pkg/session.SessionModel 구현 (yongol codegen from ssac/pkg/session/interface.yaml)"
