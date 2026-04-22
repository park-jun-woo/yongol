//ff:func feature=validate type=rule control=iteration dimension=1 topic=func-check
//ff:what XSF-62 — func spec → @call 사용

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsf62FuncSpecUsed validates XSF-62: project func spec → @call 사용 여부.
// WARNING level — project-defined func spec that is never referenced by any
// SSaC @call is dead code. YongolPkgSpecs are excluded from coverage.
func xsf62FuncSpecUsed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.ProjectFuncSpecs) == 0 {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	callRefs := g.Lookup["SSaC.callRef"]
	if callRefs == nil {
		callRefs = map[string]bool{}
	}

	var diags []diagnostic.Diagnostic
	for _, sp := range fs.ProjectFuncSpecs {
		key := sp.Package + "." + sp.Name
		if callRefs[key] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			Line:    sp.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XSF-62] func spec " + sp.Package + "." + sp.Name + " is not referenced by any SSaC @call",
			Advice:  "사용되지 않는 func " + sp.Package + "." + sp.Name + " 를 제거하거나 SSaC @call 에서 호출하세요",
		})
	}
	return diags
}
