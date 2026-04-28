//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what XNS-90 — manifest.session.backend=postgres 시 canonical DDL + sqlc 쿼리 존재 강제

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xns90SessionBackendRequiresSQLC mirrors XNC-90 for the session backend.
// Required entities are derived from ssac/pkg/session/interface.yaml:
//
//   - DDL table: fullend_sessions
//   - sqlc queries: SessionSet / SessionGet / SessionDelete
//
// The shared validateBuiltinBackend engine in xnc_90_* produces the
// diagnostic; this file only supplies the spec.
func xns90SessionBackendRequiresSQLC(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return validateBuiltinBackend(fs, backendSpec{
		Pkg:        "session",
		Cfg:        sessionCfg(fs),
		RequireDDL: "fullend_sessions",
		RequireQueries: []string{
			"SessionSet", "SessionGet", "SessionDelete",
		},
		RuleID: "XNS-90",
	})
}
