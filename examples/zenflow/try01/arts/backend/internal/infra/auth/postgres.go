//ff:type feature=infra type=model topic=auth-refresh
//ff:what postgresRefreshStore — ssac/pkg/auth.RefreshStore 구현 (yongol codegen from ssac/pkg/auth/interface.yaml)
//ff:checked llm=yongol-gen hash=d79697e1

package auth

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/park-jun-woo/ssac/pkg/auth"

	"github.com/example/zenflow_try01/internal/db"
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
// interface.yaml port: RefreshTokenInsert.
func (s *postgresRefreshStore) Create(ctx context.Context, token string, claims any, expiresAt time.Time) error {
	raw, err := auth.MarshalClaimsJSON(claims)
	if err != nil {
		return err
	}
	return s.q.RefreshTokenInsert(ctx, db.RefreshTokenInsertParams{
		TokenHash: auth.HashRefreshToken(token),
		Claims:    raw,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
}

// Consume implements one-time-use rotation as a SELECT+UPDATE pair:
//
//  1. RefreshTokenFindByHash returns claims / expires_at / revoked_at for the hash.
//  2. If the row is revoked, surface ErrRefreshTokenReused with the stored
//     claims so reuse-detection lockout can scope a family revoke.
//  3. If expired, treat as not-found.
//  4. Otherwise call RefreshTokenRevoke to mark it revoked atomically (the sqlc query
//     uses WHERE revoked_at IS NULL so concurrent Consume calls return
//     at most one winner).
//
// interface.yaml ports: RefreshTokenFindByHash, RefreshTokenRevoke.
func (s *postgresRefreshStore) Consume(ctx context.Context, token string) (json.RawMessage, error) {
	hash := auth.HashRefreshToken(token)
	row, err := s.q.RefreshTokenFindByHash(ctx, hash)
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
	if err := s.q.RefreshTokenRevoke(ctx, hash); err != nil {
		return nil, err
	}
	return row.Claims, nil
}

// Revoke marks a single refresh-token row as revoked. Idempotent — the
// sqlc query narrows to WHERE revoked_at IS NULL so a second Revoke on
// an already-revoked token is a no-op with nil error.
// interface.yaml port: RefreshTokenRevoke.
func (s *postgresRefreshStore) Revoke(ctx context.Context, token string) error {
	return s.q.RefreshTokenRevoke(ctx, auth.HashRefreshToken(token))
}

// RevokeAll revokes every active row whose stored JSONB claims contain
// every key/value in matcher. Empty matcher must be rejected before the
// DB call — unbounded revocation is a bug, never an intent.
// interface.yaml port: RefreshTokenRevokeAll.
func (s *postgresRefreshStore) RevokeAll(ctx context.Context, matcher auth.ClaimMatcher) error {
	if len(matcher) == 0 {
		return auth.ErrEmptyMatcher
	}
	raw, err := json.Marshal(matcher)
	if err != nil {
		return err
	}
	return s.q.RefreshTokenRevokeAll(ctx, raw)
}
