//ff:func feature=generate type=test-helper control=sequence
//ff:what prepareSessionBackendCases — 4케이스 테이블 생성

package prepared

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// prepareSessionBackendCases returns the canonical 4-row matrix used by
// TestPrepareSessionBackend (manifest declared O/X × SSaC used O/X).
func prepareSessionBackendCases() []prepareSessionBackendCase {
	sessionCall := ssac.ServiceFunc{
		Name:      "login",
		Sequences: []ssac.Sequence{{Type: "call", Model: "session.Put"}},
	}
	nonSession := ssac.ServiceFunc{
		Name:      "ping",
		Sequences: []ssac.Sequence{{Type: "call", Model: "auth.VerifyPassword"}},
	}
	declaredPostgres := &pmanifest.ProjectConfig{Session: &pmanifest.BuiltinBackend{Backend: "postgres"}}
	declaredMemory := &pmanifest.ProjectConfig{Session: &pmanifest.BuiltinBackend{Backend: "memory"}}
	return []prepareSessionBackendCase{
		{name: "manifest_absent_ssac_unused", wantNil: true},
		{name: "manifest_absent_ssac_used", funcs: []ssac.ServiceFunc{sessionCall}, wantBE: "memory"},
		{name: "manifest_declared_ssac_unused", manifest: declaredPostgres, funcs: []ssac.ServiceFunc{nonSession}, wantBE: "postgres"},
		{name: "manifest_declared_ssac_used", manifest: declaredMemory, funcs: []ssac.ServiceFunc{sessionCall}, wantBE: "memory"},
	}
}
