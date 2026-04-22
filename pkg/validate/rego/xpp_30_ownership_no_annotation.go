//ff:func feature=validate type=rule control=iteration dimension=1 topic=rego-structural
//ff:what XPP-30 — @ownership annotation is required when a Rego policy references resource_owner

package rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xpp30OwnershipNoAnnotation flags policies that reference `resource_owner`
// without declaring a corresponding @ownership mapping.
func xpp30OwnershipNoAnnotation(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		if !usesResourceOwner(p.Rules) {
			continue
		}
		if len(p.Ownerships) == 0 {
			diags = append(diags, diagnostic.Diagnostic{
				File:    p.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XPP-30] policy references resource_owner but declares no @ownership",
				Advice:  "Add # @ownership table.column [join_table.fk] at the top of the policy package",
			})
		}
	}
	return diags
}
