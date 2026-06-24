//ff:func feature=gen-gogin type=test control=sequence
//ff:what opIDsInDoc — nil doc / 빈 operationId / 정상 op 멤버십 집합 검증
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOpIDsInDoc(t *testing.T) {
	// nil doc → empty (non-nil) set.
	if got := opIDsInDoc(nil); got == nil || len(got) != 0 {
		t.Fatalf("opIDsInDoc(nil) = %v, want empty non-nil", got)
	}
	// nil Paths → empty set.
	if got := opIDsInDoc(&openapi3.T{}); len(got) != 0 {
		t.Fatalf("opIDsInDoc(no paths) = %v, want empty", got)
	}
	// One op with an ID is included; an op with empty OperationID is skipped.
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/a", &openapi3.PathItem{
				Get:  &openapi3.Operation{OperationID: "ListThings"},
				Post: &openapi3.Operation{OperationID: ""},
			}),
			openapi3.WithPath("/b", &openapi3.PathItem{
				Delete: &openapi3.Operation{OperationID: "DeleteThing"},
			}),
		),
	}
	got := opIDsInDoc(doc)
	if !got["ListThings"] || !got["DeleteThing"] {
		t.Fatalf("opIDsInDoc missing expected ids: %v", got)
	}
	if len(got) != 2 {
		t.Fatalf("opIDsInDoc size = %d, want 2 (empty OperationID skipped): %v", len(got), got)
	}
}
