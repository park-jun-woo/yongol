//ff:func feature=gen-gogin type=generator control=iteration dimension=2
//ff:what blockAuthzInit — OPA authz.Init(policyPath, ownerships) — DB 의존 없음

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockAuthzInit produces the authz.Init call with OwnershipMapping literals
// extracted from parsed Rego @ownership annotations. Active when any SSaC
// function uses @auth.
//
// Phase002 (ssac/purify) — ssac/pkg/authz is DB-free: the old
// authz.Init(conn, ownerships) signature has been replaced with
// authz.Init(policyPath, ownerships). OPA policy loading now drives off the
// OPA_POLICY_PATH env var (or the explicit first argument), and ownership
// lookups are performed by handler codegen via user sqlc queries (Phase003).
//
// This block therefore emits a thin initAuthz(policyPath) helper that calls
// through to authz.Init with the collected ownerships. The `conn` symbol is
// no longer threaded here.
func blockAuthzInit(fs *yongol.Fullstack) MainBlock {
	var mappings []string
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			line := fmt.Sprintf(`{Resource: %q, Table: %q, Column: %q`, om.Resource, om.Table, om.Column)
			if om.JoinTable != "" {
				line += fmt.Sprintf(`, JoinTable: %q, JoinFK: %q`, om.JoinTable, om.JoinFK)
			}
			line += `},`
			mappings = append(mappings, "\t\t"+line)
		}
	}
	factory := authzHelperInitAuthzFactory(mappings)
	lines := []string{
		`initAuthz(os.Getenv("OPA_POLICY_PATH"))`,
	}
	return MainBlock{
		Name:   "authz-init",
		Active: hasAuthSequence,
		Imports: []string{
			`"os"`,
			`"log/slog"`,
			`"github.com/park-jun-woo/ssac/pkg/authz"`,
		},
		Lines: lines,
		Funcs: []string{factory},
	}
}
