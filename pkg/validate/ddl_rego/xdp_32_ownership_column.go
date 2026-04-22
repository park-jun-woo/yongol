//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-32 — @ownership column → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp32OwnershipColumn validates XDP-32: the column referenced by a Rego
// @ownership annotation must exist in the DDL column definitions of the target
// table. Missing tables are already reported by XDP-31 and are skipped here.
func xdp32OwnershipColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()

	columnsByTable := buildDDLColumnIndex(fs)
	tables := buildDDLTableSet(fs, g)

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.Table == "" || om.Column == "" {
				continue
			}
			if !tables[om.Table] {
				continue // reported by XDP-31
			}
			key := p.File + "|" + om.Resource + "|" + om.Table + "." + om.Column
			if seen[key] {
				continue
			}
			seen[key] = true
			cols := columnsByTable[om.Table]
			if !cols[om.Column] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-32] @ownership %s — column %s.%s not found in DDL",
						om.Resource, om.Table, om.Column),
					Advice: fmt.Sprintf("Add column %s to DDL table %s (typically owner_id, user_id, etc.)", om.Column, om.Table),
				})
			}
		}
	}
	return diags
}
