//ff:func feature=validate type=util control=sequence topic=states
//ff:what statefulMethods — PathItem 에서 POST/PUT/DELETE operation tuple 반환

package ssac_statemachine

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// statefulMethods returns POST / PUT / DELETE operations on the given
// PathItem in a stable order so XSM-27 diagnostics are deterministic.
func statefulMethods(item *openapi3.PathItem) []statefulMethodOp {
	return []statefulMethodOp{
		{"POST", item.Post},
		{"PUT", item.Put},
		{"DELETE", item.Delete},
	}
}
