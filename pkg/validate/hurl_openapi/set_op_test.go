//ff:func feature=validate type=test-helper control=selection topic=hurl-openapi
//ff:what setOp — PathItem 에 HTTP method 별 operation 을 할당

package hurl_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// setOp attaches op to pi under the given HTTP method. openapi3's
// PathItem exposes dedicated fields per verb; centralising the switch
// makes fixture code small.
func setOp(pi *openapi3.PathItem, method string, op *openapi3.Operation) {
	switch strings.ToUpper(method) {
	case "GET":
		pi.Get = op
	case "POST":
		pi.Post = op
	case "PUT":
		pi.Put = op
	case "DELETE":
		pi.Delete = op
	case "PATCH":
		pi.Patch = op
	}
}
