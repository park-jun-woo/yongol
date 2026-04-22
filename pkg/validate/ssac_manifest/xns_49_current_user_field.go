//ff:func feature=validate type=rule control=iteration dimension=3 topic=config-check
//ff:what XNS-49 — currentUser.field → claims

package ssac_manifest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xns49CurrentUserField validates XNS-49: every currentUser.<field> reference
// in an SSaC sequence input must correspond to a field declared in manifest
// backend.auth.claims (Ground lookup key "Manifest.claims").
func xns49CurrentUserField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	claims := g.Lookup["Manifest.claims"]
	if claims == nil {
		// No claims configured. XNS-48 handles the "currentUser used but claims
		// missing" error; XNS-49 only flags unknown field references when
		// claims are declared.
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			for _, v := range seq.Inputs {
				if !strings.HasPrefix(v, "currentUser.") {
					continue
				}
				field := strings.TrimPrefix(v, "currentUser.")
				if claims[field] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XNS-49] references currentUser.%s but the field is not declared in backend.auth.claims", field),
					Advice:  fmt.Sprintf("manifest claims 에 필드 %s 를 추가하세요", field),
				})
			}
		}
	}
	return diags
}
