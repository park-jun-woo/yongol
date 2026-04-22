//ff:func feature=validate type=rule control=iteration dimension=3 topic=config-check
//ff:what XNS-73 — JWT @call input → claims

package ssac_manifest

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// jwtBuiltinFuncs are claims-dependent builtin functions that are generated
// under internal/auth/ when backend.auth.type is jwt. Their @call input keys
// must match manifest claim field names (not claim keys).
var jwtBuiltinFuncs = map[string]bool{
	"auth.issueToken":   true,
	"auth.verifyToken":  true,
	"auth.refreshToken": true,
}

// xns73JwtCallClaims validates XNS-73: for every SSaC @call to a JWT builtin
// function, each input key must exist in manifest backend.auth.claims field
// set (Ground Lookup["Manifest.claims"]).
func xns73JwtCallClaims(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	claims := g.Lookup["Manifest.claims"]
	if claims == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" {
				continue
			}
			if !jwtBuiltinFuncs[normalizeCallKey(seq.Model)] {
				continue
			}
			for inputKey := range seq.Inputs {
				if claims[inputKey] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  fn.FileName,
					Line:  seq.Line,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XNS-73] @call %s input %q is not a declared claim field (valid: %s)",
						seq.Model, inputKey, sortedClaimFields(claims)),
					Advice: fmt.Sprintf("manifest claims 에 %s 를 추가하거나 jwt builtin 호출에서 제거하세요", inputKey),
				})
			}
		}
	}
	return diags
}

