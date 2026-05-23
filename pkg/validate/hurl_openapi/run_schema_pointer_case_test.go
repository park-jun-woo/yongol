//ff:func feature=validate type=test-helper control=sequence topic=hurl-openapi
//ff:what runSchemaPointerCase — schema pointer 비교 테스트 공통 헬퍼

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func runSchemaPointerCase(t *testing.T, got *openapi3.Schema, wantNil bool, wantPtr *openapi3.Schema) {
	t.Helper()
	if wantNil {
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
		return
	}
	if got != wantPtr {
		t.Errorf("got %p, want %p", got, wantPtr)
	}
}
