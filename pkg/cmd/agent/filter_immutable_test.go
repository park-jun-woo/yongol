//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestFilterImmutable — immutable 파일의 diagnostic 만 제외하고 나머지 보존 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestFilterImmutable(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{File: "features.yaml", Message: "drop"},
		{File: "specs/openapi.yaml", Message: "keep1"},
		{File: "tests/x.hurl", Message: "drop"},
		{File: "db/schema.sql", Message: "keep2"},
	}
	out := filterImmutable(diags)
	if len(out) != 2 {
		t.Fatalf("filterImmutable len = %d, want 2: %+v", len(out), out)
	}
	if out[0].Message != "keep1" || out[1].Message != "keep2" {
		t.Errorf("unexpected survivors: %+v", out)
	}
}
