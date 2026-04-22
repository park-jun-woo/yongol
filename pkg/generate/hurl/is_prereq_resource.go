//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what isPrereqResource — 리소스가 auth body FK 선행 리소스인지 판정
package hurl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// isPrereqResource checks if a resource is needed as a prerequisite for auth.
func isPrereqResource(fs *yongol.Fullstack, resource string) bool {
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil || fs.OpenAPIDoc == nil {
		return false
	}
	needed := collectAuthFKResources(fs)
	for name := range needed {
		if strings.EqualFold(name, resource) {
			return true
		}
	}
	return false
}
