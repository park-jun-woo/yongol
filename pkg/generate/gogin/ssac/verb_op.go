//ff:type feature=gen-gogin type=model
//ff:what verbOp — PathItem 순회용 (HTTP verb, *Operation) 페어
package ssac

import "github.com/getkin/kin-openapi/openapi3"

type verbOp struct {
	method string
	op     *openapi3.Operation
}
