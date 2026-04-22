//ff:func feature=gen-gogin type=generator control=iteration dimension=2
//ff:what blockAuthzInit — OPA authz.Init + OwnershipMapping 리터럴 블록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockAuthzInit produces the authz.Init call with OwnershipMapping literals
// extracted from parsed Rego @ownership annotations. Active when any SSaC
// function uses @auth.
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
		`initAuthz(conn)`,
	}
	return MainBlock{
		Name:    "authz-init",
		Active:  hasAuthSequence,
		Imports: []string{
			`"database/sql"`,
			`"os"`,
			`"log/slog"`,
			`"github.com/park-jun-woo/ssac/pkg/authz"`,
		},
		Lines:   lines,
		Funcs:   []string{factory},
	}
}
