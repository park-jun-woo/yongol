//ff:func feature=gen-gogin type=util control=sequence topic=path-convert
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
		if p[i] == '{' {
			end := strings.IndexByte(p[i:], '}')
			if end > 0 {
				b.WriteByte(':')
				b.WriteString(p[i+1 : i+end])
				i += end + 1
				continue
			}
		}
		b.WriteByte(p[i])
		i++
	}
	return b.String()
}
