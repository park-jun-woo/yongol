//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what O-2 — 여러 path 에 걸쳐 case-only 로 다른 path parameter 감지

package openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o02PathParamCaseConflict validates O-2: path parameter names must be
// case-consistent across all paths in the OpenAPI document. `{id}` in one
// path and `{ID}` in another produces an ERROR because Go Gin/Echo routers
// are case-sensitive; two case variants mean two different parameters for
// the runtime while sharing intent in source — user confusion / codegen
// struct field collision.
func o02PathParamCaseConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}
	// lowercase key → original case → set of paths where it appeared
	caseMap := map[string]map[string]map[string]bool{}
	for path := range fs.OpenAPIDoc.Paths.Map() {
		for _, seg := range strings.Split(path, "/") {
			if !strings.HasPrefix(seg, "{") || !strings.HasSuffix(seg, "}") {
				continue
			}
			name := seg[1 : len(seg)-1]
			lower := strings.ToLower(name)
			if caseMap[lower] == nil {
				caseMap[lower] = map[string]map[string]bool{}
			}
			if caseMap[lower][name] == nil {
				caseMap[lower][name] = map[string]bool{}
			}
			caseMap[lower][name][path] = true
		}
	}

	var diags []diagnostic.Diagnostic
	for lower, cases := range caseMap {
		if d, ok := o02CaseConflictDiag(lower, cases); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
