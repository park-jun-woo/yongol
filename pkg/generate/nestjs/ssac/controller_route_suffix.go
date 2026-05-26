//ff:func feature=gen-nestjs type=util control=sequence
//ff:what controllerRouteSuffix — URLPath → NestJS 라우트 suffix 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// controllerRouteSuffix extracts the suffix after the base prefix.
// For example "/courses/:id/enroll" yields ":id/enroll".
func controllerRouteSuffix(plan *ir.ServicePlan) string {
	path := plan.URLPath
	if path == "" {
		return ""
	}
	path = strings.TrimPrefix(path, "/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 2 {
		return ""
	}
	return nestURLPath(parts[1])
}
