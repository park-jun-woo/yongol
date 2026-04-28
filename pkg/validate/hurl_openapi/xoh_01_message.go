//ff:func feature=validate type=rule control=sequence topic=hurl-openapi
//ff:what xoh01Message — path 존재 여부에 따라 XOH-01 진단 문구 선택

package hurl_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh01Message picks the right diagnostic text based on whether the
// path exists at all. Listing the available methods when the path
// matches turns a drift report into a near copy-pasteable fix.
func xoh01Message(e hurl.HurlEntry, segs []string, routes []apiRoute) (string, string) {
	if findPathMatch(segs, routes) < 0 {
		return "[XOH-01] " + e.Method + " " + e.Path + " — path not declared in OpenAPI",
			"Add a matching operation to openapi.yaml, or fix the hurl request path"
	}
	methods := methodsForPath(segs, routes)
	return "[XOH-01] " + e.Method + " " + e.Path + " — method not declared on this path (OpenAPI lists " + strings.Join(methods, ", ") + ")",
		"Use one of the declared methods or add " + e.Method + " to the operation"
}
