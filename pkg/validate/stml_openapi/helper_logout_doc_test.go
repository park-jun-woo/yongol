//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what logoutDoc — POST /auth/logout op 가진 OpenAPI doc 생성 (public=true면 security:[] opt-out)

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

func logoutDoc(opID string, public bool) *openapi3.T {
	var sec *openapi3.SecurityRequirements
	if public {
		s := openapi3.SecurityRequirements{}
		sec = &s
	} else {
		s := openapi3.SecurityRequirements{
			openapi3.SecurityRequirement{"bearerAuth": []string{}},
		}
		sec = &s
	}
	op := &openapi3.Operation{
		OperationID: opID,
		Security:    sec,
		Responses:   openapi3.NewResponses(),
	}
	return &openapi3.T{
		Paths: openapi3.NewPaths(openapi3.WithPath("/auth/logout", &openapi3.PathItem{Post: op})),
	}
}
