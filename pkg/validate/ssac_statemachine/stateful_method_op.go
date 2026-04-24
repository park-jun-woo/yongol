//ff:type feature=validate type=model topic=states
//ff:what statefulMethodOp — POST/PUT/DELETE 메서드와 operation 포인터 묶음

package ssac_statemachine

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// statefulMethodOp is the tuple yielded by statefulMethods so downstream
// iteration avoids switch-on-method noise.
type statefulMethodOp struct {
	method string
	op     *openapi3.Operation
}
