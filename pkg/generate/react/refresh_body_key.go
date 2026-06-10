//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what refreshBodyKey — refresh op 의 requestBody 에 refresh_field 프로퍼티가 있으면 그 키를 반환

package react

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// refreshBodyKey returns refreshField when the operation's JSON requestBody
// schema declares a property of that name — the generated refresh call then
// sends `{ <key>: <stored refresh token> }`. Empty means the op takes no
// such property (e.g. the refresh token travels in a cookie) and the call
// sends no body.
func refreshBodyKey(op *openapi3.Operation, refreshField string) string {
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return ""
	}
	for ct, mt := range op.RequestBody.Value.Content {
		if !strings.Contains(strings.ToLower(ct), "json") {
			continue
		}
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		if _, ok := mt.Schema.Value.Properties[refreshField]; ok {
			return refreshField
		}
	}
	return ""
}
