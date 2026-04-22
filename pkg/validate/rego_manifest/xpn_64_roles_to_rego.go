//ff:func feature=validate type=rule control=iteration dimension=2 topic=config-check
//ff:what XPN-64 — every manifest role must be referenced by at least one Rego allow rule

package rego_manifest

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xpn64RolesToRego validates XPN-64: every role declared in manifest
// backend.auth.roles is referenced by at least one Rego allow rule.
func xpn64RolesToRego(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	manifestRoles := g.Lookup["Manifest.roles"]
	if len(manifestRoles) == 0 {
		return nil
	}
	regoRoles := g.Lookup["Rego.roles"]

	var unused []string
	for role := range manifestRoles {
		if regoRoles[role] {
			continue
		}
		unused = append(unused, role)
	}
	sort.Strings(unused)

	var diags []diagnostic.Diagnostic
	for _, role := range unused {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[XPN-64] manifest role %q — not referenced by any Rego allow rule", role),
			Advice:  fmt.Sprintf("Add a Rego allow rule that uses role '%s', or remove it from the manifest", role),
		})
	}
	return diags
}
