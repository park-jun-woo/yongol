//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what tm54Doc — GET GetRule + PUT UpdateRule(요청 sheet_name/start_row/note) 테스트 doc

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// tm54Doc builds the edit-page OpenAPI doc shared by the TM-54/55/56 tests:
// a GET-by-id GetRule (response sheet_name, start_row) and a PUT UpdateRule
// (requestBody sheet_name, start_row, note; sheet_name/start_row required).
func tm54Doc() *openapi3.T {
	return makeDoc(map[string]*openapi3.PathItem{
		"/rules/{id}": editPath(
			"GetRule", map[string]*openapi3.SchemaRef{"sheet_name": stringProp(), "start_row": intProp()},
			"PUT", "UpdateRule", map[string]*openapi3.SchemaRef{"sheet_name": stringProp(), "start_row": intProp(), "note": stringProp()},
			[]string{"sheet_name", "start_row"},
		),
	})
}
