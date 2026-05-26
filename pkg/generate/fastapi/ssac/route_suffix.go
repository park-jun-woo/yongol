//ff:func feature=gen-fastapi type=util control=sequence
//ff:what routeSuffix — URLPath → FastAPI 라우트 suffix 추출 (prefix 제외)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// routeSuffix extracts the path portion after the base feature prefix and
// converts OpenAPI {param} to FastAPI {param} path syntax (they use the
// same syntax, so no conversion needed).
func routeSuffix(plan *ir.ServicePlan) string {
	path := plan.URLPath
	if path == "" {
		return "/"
	}
	// FastAPI uses {param} just like OpenAPI; revert any NestJS :param → {param}
	path = revertNestifyPath(path)
	// Strip the feature prefix
	path = strings.TrimPrefix(path, "/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 2 {
		return "/"
	}
	return "/" + parts[1]
}
