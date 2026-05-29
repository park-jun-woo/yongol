//ff:func feature=chain type=test control=sequence
//ff:what Chain 호출 시 존재하지 않는 operationId 에 대해 에러 반환 검증
package chain

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestChain_OperationNotFound(t *testing.T) {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info:    &openapi3.Info{Title: "t", Version: "1"},
		Paths:   openapi3.NewPaths(),
	}
	fs := &yongol.Fullstack{
		SpecsDir:   t.TempDir(),
		OpenAPIDoc: doc,
	}
	if _, err := Chain(fs, "DoesNotExist"); err == nil {
		t.Fatal("expected error for missing operationId")
	}
}
