//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprNilDocReturnsNil — Components nil 인 OpenAPI doc 은 nil 진단 반환

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func subtestTestCanonicalResponseReprNilDocReturnsNil(t *testing.T) {
	fs := buildCanonicalFS(nil, openapi3.Schemas{}, nil, nil, nil)
	fs.OpenAPIDoc.Components = nil
	if diags := canonicalResponseRepr(fs); diags != nil {
		t.Fatalf("expected nil, got %+v", diags)
	}
}
