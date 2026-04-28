package infra

// authWrapperTemplate is the printf-style template for the generated
// arts/backend/internal/infra/auth/postgres.go file. `%[1]s` is the user
// module path; `%[2]s`-`%[5]s` are the sqlc query names pulled from
// interface.yaml (Insert / FindByHash / Revoke / RevokeAll).
//
// The leading `//ff:` lines are part of the emitted file, not this file's
// own annotations — we assemble them from string fragments below so filefunc
// does not mistake them for a second header annotation on template_auth.go.
var authWrapperTemplate = authWrapperHeaderType + authWrapperHeaderWhat + `

package auth

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/park-jun-woo/ssac/pkg/auth"

	"%[1]s/internal/db"
)

// postgresRefreshStore adapts the user's sqlc Queries onto auth.RefreshStore.
// Construct via NewPostgres(queries); wire from main.go via
// auth.Init(infraauth.NewPostgres(queries)).
type postgresRefreshStore struct {
	q *db.Queries
}

// NewPostgres returns an auth.RefreshStore backed by the user's sqlc Queries.
// All refresh-token state (hash / claims / expires_at / revoked_at) lives in
// the user-owned refresh_tokens table.
func NewPostgres(q *db.Queries) auth.RefreshStore {
	return &postgresRefreshStore{q: q}
}

// Create persists a new refresh-token row. The plaintext token is hashed to
// sha256 before storage (auth.HashRefreshToken); claims is marshalled to
// JSONB so the store stays claim-schema-agnostic.
// interface.yaml port: %[2]s.
func (s *postgresRefreshStore) Create(ctx context.Context, token string, claims any, expiresAt time.Time) error {
	raw, err := auth.MarshalClaimsJSON(claims)
	if err != nil {
		return err
	}
	return s.q.%[2]s(ctx, db.%[2]sParams{
		TokenHash: auth.HashRefreshToken(token),
		Claims:    raw,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
}

// Consume implements one-time-use rotation as a SELECT+UPDATE pair:
//
//  1. %[3]s returns claims / expires_at / revoked_at for the hash.
//  2. If the row is revoked, surface ErrRefreshTokenReused with the stored
//     claims so reuse-detection lockout can scope a family revoke.
//  3. If expired, treat as not-found.
//  4. Otherwise call %[4]s to mark it revoked atomically (the sqlc query
//     uses WHERE revoked_at IS NULL so concurrent Consume calls return
//     at most one winner).
//
// interface.yaml ports: %[3]s, %[4]s.
func (s *postgresRefreshStore) Consume(ctx context.Context, token string) (json.RawMessage, error) {
	hash := auth.HashRefreshToken(token)
	row, err := s.q.%[3]s(ctx, hash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrRefreshTokenNotFound
		}
		return nil, err
	}
	if row.RevokedAt.Valid {
		return row.Claims, auth.ErrRefreshTokenReused
	}
	if row.ExpiresAt.Valid && time.Now().After(row.ExpiresAt.Time) {
		return nil, auth.ErrRefreshTokenNotFound
	}
	if err := s.q.%[4]s(ctx, hash); err != nil {
		return nil, err
	}
	return row.Claims, nil
}

// Revoke marks a single refresh-token row as revoked. Idempotent — the
// sqlc query narrows to WHERE revoked_at IS NULL so a second Revoke on
// an already-revoked token is a no-op with nil error.
// interface.yaml port: %[4]s.
func (s *postgresRefreshStore) Revoke(ctx context.Context, token string) error {
	return s.q.%[4]s(ctx, auth.HashRefreshToken(token))
}

// RevokeAll revokes every active row whose stored JSONB claims contain
// every key/value in matcher. Empty matcher must be rejected before the
// DB call — unbounded revocation is a bug, never an intent.
// interface.yaml port: %[5]s.
func (s *postgresRefreshStore) RevokeAll(ctx context.Context, matcher auth.ClaimMatcher) error {
	if len(matcher) == 0 {
		return auth.ErrEmptyMatcher
	}
	raw, err := json.Marshal(matcher)
	if err != nil {
		return err
	}
	return s.q.%[5]s(ctx, raw)
}
`

// authWrapperHeader* assemble the emitted file's `//ff:` annotations from
// string literals that don't start with `//` so this file's own filefunc
// scan does not see them as duplicate top-of-file annotations.
var authWrapperHeaderType = "//" + "ff:type feature=infra type=model topic=auth-refresh\n"
var authWrapperHeaderWhat = "//" + "ff:what postgresRefreshStore — ssac/pkg/auth.RefreshStore 구현 (yongol codegen from ssac/pkg/auth/interface.yaml)"
