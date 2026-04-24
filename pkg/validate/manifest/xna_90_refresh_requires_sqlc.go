//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XNA-90 — backend.auth 구성 시 refresh_tokens DDL + sqlc 쿼리 존재 강제

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xna90RefreshRequiresSQLC mirrors XNC-90 for ssac/pkg/auth's refresh
// rotation subsystem. Unlike cache/session/queue, auth does not have a
// separate `memory` backend toggle — presence of `backend.auth` in the
// manifest implies the user intends to issue refresh tokens, so the
// refresh_tokens schema must be present.
//
// Required entities (from ssac/pkg/auth/interface.yaml):
//
//   - DDL table: refresh_tokens
//   - sqlc queries: RefreshTokenInsert / RefreshTokenFindByHash /
//     RefreshTokenRevoke / RefreshTokenRevokeAll
//
// LoginLookup is also declared in interface.yaml but is NOT enforced by
// this rule — it is a user-facing port that lives alongside the user's
// own `users` table (the table name is project-specific, so there is no
// canonical `fullend_users`).
func xna90RefreshRequiresSQLC(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	// "postgres" is assumed whenever backend.auth is configured — there is
	// no memory refresh backend in ssac today (see interface.yaml
	// `when: manifest.backend.auth.refresh.enabled`). Treat the presence
	// flag as equivalent to "postgres" so validateBuiltinBackend's gate
	// fires correctly.
	return validateBuiltinBackend(fs, backendSpec{
		Pkg:        "auth",
		Cfg:        builtinBackend{Present: true, Backend: "postgres"},
		RequireDDL: "refresh_tokens",
		RequireQueries: []string{
			"RefreshTokenInsert",
			"RefreshTokenFindByHash",
			"RefreshTokenRevoke",
			"RefreshTokenRevokeAll",
		},
		RuleID: "XNA-90",
	})
}
