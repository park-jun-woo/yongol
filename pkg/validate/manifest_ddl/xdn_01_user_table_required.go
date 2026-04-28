//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XDN-01 — backend.auth 활성 시 user_table 필수 (없으면 ERROR)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdn01UserTableRequired enforces that any manifest with an active auth
// block (auth.type != "none") names the DDL table that backs the user
// rows from which JWT claims derive. The legacy "infer from db/users.sql"
// convention breaks under non-standard naming (`accounts`, `members`) or
// multi-tenant auth schemes (admin / customer split); making the field
// explicit costs one boilerplate line and removes the ambiguity entirely.
//
// Skipped when auth is absent or type is "none" — the field has no
// meaning without an active auth subsystem.
func xdn01UserTableRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !isAuthActive(fs) {
		return nil
	}
	auth := fs.Manifest.Backend.Auth
	if auth.UserTable != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Line:    0,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XDN-01] backend.auth.user_table is required when auth is active",
		Advice: "Add `user_table: <table-name>` under backend.auth — the DDL table that holds user " +
			"rows backing the JWT claims (e.g. `user_table: users`).",
	}}
}
