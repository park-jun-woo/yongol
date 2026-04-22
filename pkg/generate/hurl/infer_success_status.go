//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what inferSuccessStatus — operation responses에서 첫 2xx 상태 코드 추출
package hurl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// inferSuccessStatus returns the first 2xx status code from the operation's responses.
func inferSuccessStatus(op *openapi3.Operation) int {
	if op == nil || op.Responses == nil {
		return 200
	}
	for code := range op.Responses.Map() {
		if strings.HasPrefix(code, "2") {
			return statusCodeToInt(code)
		}
	}
	return 200
}
