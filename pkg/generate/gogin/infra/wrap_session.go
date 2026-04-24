//ff:func feature=gen-gogin type=generator control=sequence topic=session
//ff:what emitSessionWrapper — ssac/pkg/session.SessionModel 을 사용자 sqlc Queries 로 구현하는 adapter 를 emit

package infra

import (
	"bytes"
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitSessionWrapper writes `arts/backend/internal/infra/session/postgres.go`.
// The shape mirrors wrap_cache — session and cache share the same
// TTL-keyed interface (Set/Get/Delete), the only differences being the
// fullend_sessions table, the Session* sqlc names, and the SessionModel
// interface. The same `any → JSON`, `[]byte → string`, `ttl → expires_at`
// glue applies.
func emitSessionWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	setPort := portByName(active, "SessionSet")
	getPort := portByName(active, "SessionGet")
	delPort := portByName(active, "SessionDelete")
	if setPort == nil || getPort == nil || delPort == nil {
		return fmt.Errorf("session: interface.yaml missing one of SessionSet/SessionGet/SessionDelete (active ports: %d)", len(active))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, `//ff:type feature=infra type=model topic=session
//ff:what postgresSession — ssac/pkg/session.SessionModel 구현 (yongol codegen from ssac/pkg/session/interface.yaml)

package session

import (
	"context"
	"encoding/json"
	"time"

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
		return "", err
	}
	return string(raw), nil
}

// Delete forwards to the sqlc query. interface.yaml port: %[4]s.
func (s *postgresSession) Delete(ctx context.Context, key string) error {
	return s.q.%[4]s(ctx, key)
}
`, modulePath, setPort.Name, getPort.Name, delPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
