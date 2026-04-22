//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-31 — @ownership table → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp31OwnershipTable validates XDP-31: every table referenced by a Rego
// @ownership annotation must exist in the DDL.
func xdp31OwnershipTable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	tables := g.Lookup["DDL.table"]
	if tables == nil {
		// Missing Ground is a pipeline bug upstream; XDP-31 relies solely on Ground.
		return nil
	}

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.Table == "" {
				continue
			}
			key := p.File + "|" + om.Resource + "|" + om.Table
			if seen[key] {
				continue
			}
			seen[key] = true
			if !tables[om.Table] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-31] @ownership %s — table %q not found in DDL",
						om.Resource, om.Table),
					Advice: fmt.Sprintf("Define table %s in the DDL, or remove it from the Rego @ownership annotation", om.Table),
				})
			}
		}
	}
	return diags
}
