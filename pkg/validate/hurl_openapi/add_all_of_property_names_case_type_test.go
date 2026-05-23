//ff:type feature=validate type=model topic=hurl-openapi
//ff:what TestAddAllOfPropertyNamesCase — table-driven 테스트 케이스 구조체

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

type TestAddAllOfPropertyNamesCase struct {
	name     string
	allOf    openapi3.SchemaRefs
	existing map[string]struct{}
	want     map[string]struct{}
}
