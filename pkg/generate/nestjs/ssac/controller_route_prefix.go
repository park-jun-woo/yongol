//ff:func feature=gen-nestjs type=util control=sequence
//ff:what controllerRoutePrefix — URLPath → NestJS Controller 라우트 prefix 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// controllerRoutePrefix extracts the base route prefix from the plan's URLPath.
// For example "/courses/:id" yields "courses".
func controllerRoutePrefix(plan *ir.ServicePlan) string {
	path := plan.URLPath
	if path == "" {
		return lcFirst(plan.Feature)
	}
	path = strings.TrimPrefix(path, "/")
	parts := strings.SplitN(path, "/", 2)
	return parts[0]
}
