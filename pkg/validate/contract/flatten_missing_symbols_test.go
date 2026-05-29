//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestFlattenMissingSymbols — missingSymbols 3 범주를 카테고리 순서 Diagnostic 으로 직렬화 검증

package contract

import (
	"strings"
	"testing"
)

func TestFlattenMissingSymbols(t *testing.T) {
	ms := missingSymbols{
		Queries: []string{"GoneQuery"},
		Calls:   []string{"billing.Gone"},
		Fields:  []string{"u.Gone"},
	}
	diags := flattenMissingSymbols("svc/foo.go", ms)
	if len(diags) != 3 {
		t.Fatalf("expected 3 diags, got %d", len(diags))
	}
	// Category order: queries → calls → fields.
	if !strings.Contains(diags[0].Message, "GoneQuery") {
		t.Errorf("diag[0] = %q", diags[0].Message)
	}
	if !strings.Contains(diags[1].Message, "billing.Gone") {
		t.Errorf("diag[1] = %q", diags[1].Message)
	}
	if !strings.Contains(diags[2].Message, "Gone") {
		t.Errorf("diag[2] = %q", diags[2].Message)
	}

	t.Run("empty → no diags", func(t *testing.T) {
		if d := flattenMissingSymbols("x.go", missingSymbols{}); len(d) != 0 {
			t.Errorf("expected no diags, got %d", len(d))
		}
	})
}
