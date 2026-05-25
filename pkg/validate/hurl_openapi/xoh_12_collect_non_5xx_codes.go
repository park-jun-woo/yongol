//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what collectNon5xxCodes — responses 에서 5xx 를 제외한 sorted status code 목록 반환

package hurl_openapi

import (
	"sort"
	"strings"
)

func collectNon5xxCodes(responses map[string]bool) []string {
	var codes []string
	for code := range responses {
		if !strings.HasPrefix(code, "5") {
			codes = append(codes, code)
		}
	}
	sort.Strings(codes)
	return codes
}
