//ff:func feature=validate type=rule control=iteration dimension=2 topic=config-check
//ff:what XNP-53 — every Rego input.claims reference must be declared in manifest claims

package rego_manifest

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnp53InputClaimsValues validates XNP-53: every `input.claims.<key>` reference
// in Rego policies corresponds to a JWT claim key declared in the manifest's
// backend.auth.claims.
func xnp53InputClaimsValues(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	claimKeys := g.Lookup["Manifest.claims.keys"]
	if len(claimKeys) == 0 {
		return nil
	}

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, ref := range p.ClaimsRefs {
			if claimKeys[ref] {
				continue
			}
			key := p.File + "|" + ref
			if seen[key] {
				continue
			}
			seen[key] = true
			diags = append(diags, diagnostic.Diagnostic{
				File:    p.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XNP-53] Rego input.claims.%s — not declared in manifest backend.auth.claims", ref),
				Advice:  fmt.Sprintf("Add field '%s' to manifest backend.auth.claims, or remove the input.claims.%s reference from Rego", ref, ref),
			})
		}
	}
	return diags
}
