//ff:func feature=gen-gogin type=util control=sequence topic=auth-refresh
//ff:what authWrapperMethodHeader — infra/auth postgres adapter 파일별 //ff:func + //ff:what 헤더 조립

package infra

// authWrapperTypeTemplate emits postgres.go with the type only.
// %[1]s = modulePath.
var authWrapperTypeTemplate = authWrapperHeaderType + authWrapperHeaderWhat + `

package auth

import "%[1]s/internal/db"

// postgresRefreshStore adapts the user's sqlc Queries onto auth.RefreshStore.
// Construct via NewPostgres(queries); wire from main.go via
// auth.Init(infraauth.NewPostgres(queries)).
type postgresRefreshStore struct {
	q *db.Queries
}
`

// authWrapperNewPostgresTemplate emits postgres_new.go with the constructor.
// %[1]s = modulePath.
var authWrapperNewPostgresTemplate = authWrapperNewHeader + `

package auth

import (
	"github.com/park-jun-woo/ssac/pkg/auth"

	"%[1]s/internal/db"
)

// NewPostgres returns an auth.RefreshStore backed by the user's sqlc Queries.
// All refresh-token state (hash / claims / expires_at / revoked_at) lives in
// the user-owned refresh_tokens table.
func NewPostgres(q *db.Queries) auth.RefreshStore {
	return &postgresRefreshStore{q: q}
}
`

// authWrapperCreateTemplate emits postgres_create.go.
// %[1]s = modulePath, %[2]s = InsertPort.
var authWrapperCreateTemplate = authWrapperMethodHeader("Create", "postgresRefreshStore — Create: refresh-token 행 생성") + `

package auth

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/park-jun-woo/ssac/pkg/auth"

	"%[1]s/internal/db"
)

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
`

// authWrapperConsumeTemplate emits postgres_consume.go.
// %[1]s = modulePath, %[2]s = FindByHashPort, %[3]s = RevokePort.
var authWrapperConsumeTemplate = authWrapperMethodHeader("Consume", "postgresRefreshStore — Consume: one-time-use rotation (SELECT+UPDATE)") + `

package auth

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/park-jun-woo/ssac/pkg/auth"
)

// Consume implements one-time-use rotation as a SELECT+UPDATE pair.
// interface.yaml ports: %[2]s, %[3]s.
func (s *postgresRefreshStore) Consume(ctx context.Context, token string) (json.RawMessage, error) {
	hash := auth.HashRefreshToken(token)
	row, err := s.q.%[2]s(ctx, hash)
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
	if err := s.q.%[3]s(ctx, hash); err != nil {
		return nil, err
	}
	return row.Claims, nil
}
`

// authWrapperRevokeTemplate emits postgres_revoke.go.
// %[1]s = modulePath, %[2]s = RevokePort.
var authWrapperRevokeTemplate = authWrapperMethodHeader("Revoke", "postgresRefreshStore — Revoke: 단일 refresh-token 무효화") + `

package auth

import (
	"context"

	"github.com/park-jun-woo/ssac/pkg/auth"
)

// Revoke marks a single refresh-token row as revoked. Idempotent.
// interface.yaml port: %[2]s.
func (s *postgresRefreshStore) Revoke(ctx context.Context, token string) error {
	return s.q.%[2]s(ctx, auth.HashRefreshToken(token))
}
`

// authWrapperRevokeAllTemplate emits postgres_revoke_all.go.
// %[1]s = modulePath, %[2]s = RevokeAllPort.
var authWrapperRevokeAllTemplate = authWrapperMethodHeader("RevokeAll", "postgresRefreshStore — RevokeAll: matcher 기반 일괄 무효화") + `

package auth

import (
	"context"
	"encoding/json"

	"github.com/park-jun-woo/ssac/pkg/auth"
)

// RevokeAll revokes every active row whose stored JSONB claims contain
// every key/value in matcher. Empty matcher must be rejected.
// interface.yaml port: %[2]s.
func (s *postgresRefreshStore) RevokeAll(ctx context.Context, matcher auth.ClaimMatcher) error {
	if len(matcher) == 0 {
		return auth.ErrEmptyMatcher
	}
	raw, err := json.Marshal(matcher)
	if err != nil {
		return err
	}
	return s.q.%[2]s(ctx, raw)
}
`

// authWrapperHeader* assemble the emitted file's annotations.
var authWrapperHeaderType = "//" + "ff:type feature=infra type=model topic=auth-refresh\n"
var authWrapperHeaderWhat = "//" + "ff:what postgresRefreshStore — ssac/pkg/auth.RefreshStore 구현 (yongol codegen from ssac/pkg/auth/interface.yaml)"
var authWrapperNewHeader = "//" + "ff:func feature=infra type=accessor control=sequence topic=auth-refresh\n" +
	"//" + "ff:what NewPostgres — postgresRefreshStore 생성자 (auth.RefreshStore 반환)"

// authWrapperMethodHeader returns the //ff:func + //ff:what header for a method file.
func authWrapperMethodHeader(method, what string) string {
	return "//" + "ff:func feature=infra type=accessor control=sequence topic=auth-refresh\n" +
		"//" + "ff:what " + what
}
