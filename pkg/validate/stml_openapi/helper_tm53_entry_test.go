//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what tm53Entry — TM-53 테스트용 다양한 타입 필드를 가진 GET operationEntry 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// tm53Entry builds a GET operationEntry whose 2xx response carries string,
// boolean, integer, object, and array fields for the TM-53 cases.
func tm53Entry() operationEntry {
	item := getOp("GetThing", nil, map[string]*openapi3.SchemaRef{
		"title":  stringProp(),
		"avatar": stringProp(),
		"active": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"boolean"}}},
		"count":  intProp(),
		"meta":   objectProp(),
		"tags":   arrayProp("string"),
	})
	return operationEntry{method: "GET", op: item.Get}
}
