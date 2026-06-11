//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what tm56Doc — GET GetRule + PATCH PatchRule(required 가변) 테스트 doc

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// tm56Doc builds a GET GetRule + PATCH PatchRule (requestBody sheet_name,
// start_row) doc; required selects which requestBody fields are required.
func tm56Doc(required []string) *openapi3.T {
	return makeDoc(map[string]*openapi3.PathItem{
		"/rules/{id}": editPath(
			"GetRule", map[string]*openapi3.SchemaRef{"sheet_name": stringProp()},
			"PATCH", "PatchRule", map[string]*openapi3.SchemaRef{"sheet_name": stringProp(), "start_row": intProp()},
			required,
		),
	})
}
