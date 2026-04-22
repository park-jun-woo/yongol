//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what substitutePathParams — OpenAPI {param}을 hurl {{resource_param}} 변수로 치환
package hurl

import (
	"regexp"
	"strings"
)

var pathParamRe = regexp.MustCompile(`\{([^}]+)\}`)

// substitutePathParams replaces {id} with {{resource_id}} in OpenAPI paths.
func substitutePathParams(path string, capturedVars map[string]bool) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if !pathParamRe.MatchString(part) {
			continue
		}
		parts[i] = resolveParamVar(parts, i)
	}
	return strings.Join(parts, "/")
}
