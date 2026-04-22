//ff:func feature=validate type=rule control=iteration dimension=2 topic=config-check
//ff:what XNP-63 — Rego role → Manifest roles

package rego_manifest

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnp63RoleManifest validates XNP-63: every role value referenced by a Rego
// allow rule must be declared in manifest backend.auth.roles.
func xnp63RoleManifest(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	roles := g.Lookup["Manifest.roles"]
	manifestDeclared := len(roles) > 0

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, r := range p.Rules {
			if !r.UsesRole || r.RoleValue == "" {
				continue
			}
			key := p.File + "|" + r.RoleValue
			if seen[key] {
				continue
			}
			seen[key] = true
			if !manifestDeclared {
				diags = append(diags, diagnostic.Diagnostic{
					File:    p.File,
					Line:    r.SourceLine,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XNP-63] Rego uses role %q but manifest backend.auth.roles is not declared", r.RoleValue),
					Advice:  fmt.Sprintf("manifest backend.auth.roles 에 [%s, ...] 를 선언하세요", r.RoleValue),
				})
				continue
			}
			if roles[r.RoleValue] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    p.File,
				Line:    r.SourceLine,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XNP-63] Rego role %q — not declared in manifest backend.auth.roles", r.RoleValue),
				Advice:  fmt.Sprintf("manifest backend.auth.roles 에 '%s' 를 추가하세요", r.RoleValue),
			})
		}
	}
	return diags
}
