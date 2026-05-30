//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — o06WalkOperation 에 nil op 직접 호출 시 acc 불변

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO06_Case12_WalkNilOp(t *testing.T) {
	visited := map[*openapi3.Schema]bool{}
	if got := o06WalkOperation(nil, visited, nil); len(got) != 0 {
		t.Fatalf("expected empty acc, got %+v", got)
	}
}
