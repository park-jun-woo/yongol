//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=path-convert
//ff:what openAPIPathToGin — OpenAPI "/x/{id}" → gin "/x/:id" 경로 변환

package boot

import "strings"

// openAPIPathToGin rewrites OpenAPI path templates ({name}) into gin route
// syntax (:name). Gin exposes this form via c.FullPath() so it is the key
// used by the runtime OverrideBodyLimit lookup.
func openAPIPathToGin(p string) string {
	var b strings.Builder
	b.Grow(len(p))
	i := 0
	for i < len(p) {
		i = openAPIPathStep(p, i, &b)
	}
	return b.String()
}
