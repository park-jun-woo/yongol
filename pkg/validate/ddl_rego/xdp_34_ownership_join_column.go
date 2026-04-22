//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-34 — @ownership via join column → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp34OwnershipJoinColumn validates XDP-34: the join FK column in a Rego
// @ownership via clause must exist in the DDL column definitions of the join
// table. Missing join tables are already reported by XDP-33 and are skipped here.
func xdp34OwnershipJoinColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	tables := buildDDLTableSet(fs, g)
	columnsByTable := buildDDLColumnIndex(fs)

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.JoinTable == "" || om.JoinFK == "" {
				continue
			}
			if !tables[om.JoinTable] {
				continue // reported by XDP-33
			}
			key := p.File + "|" + om.Resource + "|" + om.JoinTable + "." + om.JoinFK
			if seen[key] {
				continue
			}
			seen[key] = true
			cols := columnsByTable[om.JoinTable]
			if !cols[om.JoinFK] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-34] @ownership %s via — join column %s.%s not found in DDL",
						om.Resource, om.JoinTable, om.JoinFK),
					Advice: fmt.Sprintf("Add column %s to DDL join table %s", om.JoinFK, om.JoinTable),
				})
			}
		}
	}
	return diags
}
