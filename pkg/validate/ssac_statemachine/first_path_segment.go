//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what firstPathSegment — URL path 의 첫 비-중괄호 세그먼트 추출

package ssac_statemachine

import (
	"strings"
)

// firstPathSegment returns the first non-empty segment of a URL path, skipping
// leading slashes and curly-brace parameters.
func firstPathSegment(path string) string {
	for _, part := range strings.Split(path, "/") {
		if part == "" {
			continue
		}
		if strings.HasPrefix(part, "{") {
			continue
		}
		return part
	}
	return ""
}
