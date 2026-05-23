//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what runDiagCodeCase — diagnostic 개수 및 메시지 코드 검증 공통 헬퍼

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runDiagCodeCase(t *testing.T, diags []diagnostic.Diagnostic, wantCount int, code string) {
	t.Helper()
	if len(diags) != wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), wantCount, diags)
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, code) {
			t.Errorf("expected %s in message, got %q", code, d.Message)
		}
	}
}
