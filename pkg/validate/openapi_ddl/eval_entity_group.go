//ff:func feature=validate type=rule control=iteration dimension=1 topic=openapi-ddl
//ff:what evalEntityGroup — 한 엔티티의 응답 표현 집합을 평가해 XDO-11(분기)/XDO-12(inline drift) 진단 생성

package openapi_ddl

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// evalEntityGroup evaluates one entity's response representations. >1 distinct
// shape ⇒ XDO-11 (ERROR, divergent contract) as a single diagnostic listing the
// operations. Otherwise each inline operation ⇒ XDO-12 (WARNING, drift risk).
func evalEntityGroup(comp string, reprs []entityRepr) []diagnostic.Diagnostic {
	sort.Slice(reprs, func(i, j int) bool { return reprs[i].opID < reprs[j].opID })

	shapes := make(map[string]bool, len(reprs))
	for _, r := range reprs {
		shapes[r.shapeKey] = true
	}

	if len(shapes) > 1 {
		parts := make([]string, 0, len(reprs))
		for _, r := range reprs {
			parts = append(parts, fmt.Sprintf("%s=%s", r.opID, r.shapeKey))
		}
		return []diagnostic.Diagnostic{{
			File:    "api/openapi.yaml",
			Line:    reprs[0].line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XDO-11] entity %s is returned with divergent representations across operations: %s", comp, strings.Join(parts, "; ")),
			Advice:  fmt.Sprintf("A resource has exactly one representation. Make every 2xx response that returns %s $ref the same component #/components/schemas/%s", comp, comp),
		}}
	}

	var diags []diagnostic.Diagnostic
	for _, r := range reprs {
		if !strings.HasPrefix(r.shapeKey, "inline:") {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        "api/openapi.yaml",
			Line:        r.line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelWarning,
			Message:     fmt.Sprintf("[XDO-12] operation %s returns entity %s inline instead of sharing a component", r.opID, comp),
			Advice:      fmt.Sprintf("$ref #/components/schemas/%s so the representation cannot drift across operations", comp),
			OperationID: r.opID,
		})
	}
	return diags
}
