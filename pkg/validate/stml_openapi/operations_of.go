//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what operationsOf — PathItem의 GET/POST/PUT/DELETE/PATCH operation 슬라이스 반환

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// operationsOf returns the standard CRUD operations of a path item in a fixed
// order. Nil entries (methods not defined) are preserved so callers can decide
// how to handle them.
func operationsOf(item *openapi3.PathItem) []*openapi3.Operation {
	return []*openapi3.Operation{item.Get, item.Post, item.Put, item.Delete, item.Patch}
}
