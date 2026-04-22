//ff:func feature=validate type=rule control=iteration dimension=2 topic=config-check
//ff:what XPN-54 — Manifest claims → Rego 참조

package rego_manifest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xpn54ClaimsToRego validates XPN-54: every claim key declared in manifest
// backend.auth.claims is actually referenced somewhere. Exception: besides
// Rego (`input.claims.<key>`), the key is also accepted when consumed by
// middleware (Lookup["Middleware.claims"]) or exposed in an OpenAPI response
// schema (Schemas["OpenAPI.response.*"]).
func xpn54ClaimsToRego(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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

	regoRefs := g.Lookup["Rego.claims"]
	middlewareRefs := g.Lookup["Middleware.claims"]

	// responseFields merges every OpenAPI response schema's field names into a
	// single lookup set. g.Schemas values are []string (field name lists from
	// parser/ground/populate_schemas.go; duplicates may occur across operations).
	// Duplicate keys inserted into the map are harmless — the zero-value bool
	// is overwritten by true, semantics unchanged.
	responseFields := make(map[string]bool)
	for k, fields := range g.Schemas {
		if !strings.HasPrefix(k, "OpenAPI.response.") {
			continue
		}
		for _, f := range fields {
			responseFields[f] = true
		}
	}

	var diags []diagnostic.Diagnostic
	for key := range claimKeys {
		if regoRefs[key] || middlewareRefs[key] || responseFields[key] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[XPN-54] manifest claim %q — not referenced by Rego, middleware, or OpenAPI response", key),
			Advice:  fmt.Sprintf("manifest claim '%s' 를 Rego/middleware/response 중 하나에서 사용하거나 manifest 에서 제거하세요", key),
		})
	}
	return diags
}
