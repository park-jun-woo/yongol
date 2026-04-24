//ff:func feature=gen-gogin type=generator control=sequence topic=auth-refresh
//ff:what emitAuthWrapper — ssac/pkg/auth.RefreshStore 를 사용자 sqlc Queries 로 구현하는 adapter 를 emit

package infra

import (
	"bytes"
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitAuthWrapper writes `arts/backend/internal/infra/auth/postgres.go` — a
// postgres adapter that satisfies ssac/pkg/auth.RefreshStore by forwarding
// each method onto the user's sqlc Queries declared in
// ssac/pkg/auth/interface.yaml.
//
// Mapping (RefreshStore method → interface.yaml port):
//
//   Create       → RefreshTokenInsert
//   Consume      → RefreshTokenFindByHash + RefreshTokenRevoke
//                  (SELECT active row, validate, then UPDATE revoked_at —
//                   this is the only RefreshStore method that needs two
//                   queries; every other method is a single forwarder.)
//   Revoke       → RefreshTokenRevoke
//   RevokeAll    → RefreshTokenRevokeAll
//
// Type-translation glue:
//
//   - RefreshStore.Create takes `claims any`; the wrapper json.Marshals
//     it so the sqlc param stays []byte (stored as JSONB).
//   - RefreshStore.Consume plaintext token → sha256 hash before every
//     DB touch via auth.HashRefreshToken. Revoked rows surface as
//     ErrRefreshTokenReused together with their claims so reuse-detection
//     lockout can inspect the claim set.
//   - expires_at / revoked_at are pgtype.Timestamptz for pgx/v5 params;
//     the wrapper converts to/from time.Time at the boundary.
//
// LoginLookup is declared in interface.yaml `when: always` but is NOT
// implemented by this adapter — it is consumed directly by user SSaC
// @call auth.Login handlers (Phase003 handler codegen), not by the
// RefreshStore interface.
func emitAuthWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	insertPort := portByName(active, "RefreshTokenInsert")
	findPort := portByName(active, "RefreshTokenFindByHash")
	revokePort := portByName(active, "RefreshTokenRevoke")
	revokeAllPort := portByName(active, "RefreshTokenRevokeAll")
	if insertPort == nil || findPort == nil || revokePort == nil || revokeAllPort == nil {
		return fmt.Errorf("auth: interface.yaml missing one of RefreshTokenInsert/FindByHash/Revoke/RevokeAll (active ports: %d)", len(active))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, `//ff:type feature=infra type=model topic=auth-refresh
//ff:what postgresRefreshStore — ssac/pkg/auth.RefreshStore 구현 (yongol codegen from ssac/pkg/auth/interface.yaml)

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
`, modulePath, insertPort.Name, findPort.Name, revokePort.Name, revokeAllPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
