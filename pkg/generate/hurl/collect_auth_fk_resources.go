//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what collectAuthFKResources — auth endpoint body에서 _id 접미사 FK 리소스 수집
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectAuthFKResources collects resource names from _id fields in auth request bodies.
func collectAuthFKResources(fs *yongol.Fullstack) map[string]bool {
	needed := map[string]bool{}
	for _, pathItem := range fs.OpenAPIDoc.Paths.Map() {
		addAuthFKFields(pathItem, needed)
	}
	return needed
}
