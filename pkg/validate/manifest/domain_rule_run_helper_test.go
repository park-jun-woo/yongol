//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runDomainRuleCase — 도메인 규칙 1건 실행 후 진단 개수·레벨·메시지 prefix 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runDomainRuleCase runs one domains-block rule against a case and asserts the
// diagnostic count, plus the level and message prefix of every diagnostic.
func runDomainRuleCase(t *testing.T, fn func(*yongol.Fullstack) []diagnostic.Diagnostic, c domainRuleCase, level diagnostic.Level, prefix string) {
	t.Helper()
	diags := fn(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != level {
			t.Errorf("level = %q, want %q", d.Level, level)
		}
		if !strings.Contains(d.Message, prefix) {
			t.Errorf("message missing %s: %q", prefix, d.Message)
		}
	}
}
