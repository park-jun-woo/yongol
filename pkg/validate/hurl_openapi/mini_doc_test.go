//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — checkEntryURLMethod / xoh12 / xoh13 분기 직접 호출
package hurl_openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func miniDoc() *openapi3.T {
	op := openapi3.NewOperation()
	op.OperationID = "ListWidgets"
	op.Responses = openapi3.NewResponses()
	op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("ok")})
	op.Responses.Set("404", &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("missing")})
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/widgets", &openapi3.PathItem{Get: op})
	return doc
}
