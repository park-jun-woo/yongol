//ff:func feature=ground type=test-helper control=sequence
//ff:what makeDomainDoc — 도메인 식별용 한 경로 GET op + 보안 스킴을 가진 OpenAPI doc 생성

package ground

import "github.com/getkin/kin-openapi/openapi3"

// makeDomainDoc builds a one-path OpenAPI doc with a GET op and one security
// scheme so the merge test can give each domain a distinct identity.
func makeDomainDoc(path, opID, scheme string) *openapi3.T {
	p := openapi3.NewPaths(openapi3.WithPath(path, &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: opID},
	}))
	return &openapi3.T{
		Paths: p,
		Components: &openapi3.Components{
			SecuritySchemes: openapi3.SecuritySchemes{
				scheme: &openapi3.SecuritySchemeRef{Value: &openapi3.SecurityScheme{Type: "http", Scheme: "bearer"}},
			},
		},
	}
}
