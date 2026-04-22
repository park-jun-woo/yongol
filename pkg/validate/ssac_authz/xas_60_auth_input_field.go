//ff:func feature=validate type=rule control=iteration dimension=3 topic=authz-check
//ff:what XAS-60 — @auth input field → CheckRequest

package ssac_authz

import (
	"strconv"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xas60AuthInputField validates XAS-60: every @auth sequence input key must
// exist as a field in the built-in authz.CheckRequest struct. Field names are
// loaded into Ground.Lookup["Authz.checkRequest"] by populateAuthz; when the
// manifest declares a custom authz package the set is empty and this rule
// skips (we cannot introspect user-provided CheckRequest structs).
func xas60AuthInputField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	fields := g.Lookup["Authz.checkRequest"]
	if len(fields) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			for key := range seq.Inputs {
				if fields[key] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[XAS-60] @auth input field " + strconv.Quote(key) + " missing in authz.CheckRequest",
					Advice:  "Use only fields defined in authz.CheckRequest, or register a custom authz package in the manifest",
				})
			}
		}
	}
	return diags
}
