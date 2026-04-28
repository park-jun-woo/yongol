//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-statemachine
//ff:what normPath — `{foo}` / `{{var}}` / 숫자 세그먼트를 `:param` 으로 정규화

package hurl_statemachine

import (
	"regexp"
	"strings"
)

// normPath collapses `{foo}` / `{{var}}` / numeric segments into
// `:param` so hurl paths and OpenAPI paths share one canonical form.
// openapiRe is the regex to match OpenAPI-style `{var}` placeholders;
// the second arg is reserved for future tuning.
func normPath(path string, openapiRe, _ *regexp.Regexp) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var out []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		out = append(out, normPathSegment(p, openapiRe))
	}
	return "/" + strings.Join(out, "/")
}
