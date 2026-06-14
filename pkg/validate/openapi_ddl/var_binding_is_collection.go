//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what varBindingIsCollection — @response var 가 Page/Cursor/슬라이스로 바인딩됐는지 판정

package openapi_ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// varBindingIsCollection reports whether varName was bound (via @get/@post/@call
// Result) as a paginated wrapper (Page/Cursor) or a slice. Ground's
// SSaC.var.* type map records only the concrete element type (parse_result.go
// stores Result.Type = "Gig" for "Page[Gig]"), so the wrapper information is
// only recoverable from the binding sequence itself — this lets canonical
// grouping exclude list/paginated responses that resemble single resources.
func varBindingIsCollection(fn *ssac.ServiceFunc, varName string) bool {
	for _, seq := range fn.Sequences {
		if seq.Result == nil || seq.Result.Var != varName {
			continue
		}
		if seq.Result.Wrapper != "" {
			return true
		}
		return strings.HasPrefix(seq.Result.Type, "[]")
	}
	return false
}
