//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what isPrereqResource — 리소스가 auth body FK 선행 리소스인지 판정
package hurl

import (
	"strings"
)

// isPrereqResource checks if a resource is needed as a prerequisite for
// auth (i.e. referenced via <resource>_id in a signup/login request body).
// Uses ctx.authOpIDs (shape-detected) to decide which ops are auth.
func isPrereqResource(ctx *scenarioCtx, resource string) bool {
	fs := ctx.fs
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil || fs.OpenAPIDoc == nil {
		return false
	}
	needed := collectAuthFKResources(ctx)
	for name := range needed {
		if strings.EqualFold(name, resource) {
			return true
		}
	}
	return false
}
