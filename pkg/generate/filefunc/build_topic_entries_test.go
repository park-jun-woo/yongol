//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestBuildTopicEntries — 고정 topic 맵 반환 검증

package filefunc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildTopicEntries(t *testing.T) {
	got := buildTopicEntries(&yongol.Fullstack{})
	want := []string{
		"auth-check", "auth-refresh", "dos-guard", "error-envelope",
		"error-mapping", "observability", "pagination", "pointer-helper",
		"publish", "rate-limit", "request-binding", "request-id",
		"response-serialize", "security-headers", "state-transition",
		"subscribe", "transaction-boundary",
	}
	if len(got) != len(want) {
		t.Errorf("expected %d topics, got %d: %v", len(want), len(got), got)
	}
	for _, k := range want {
		if desc, ok := got[k]; !ok || desc == "" {
			t.Errorf("missing or empty topic %q: %q", k, desc)
		}
	}
	// fs is accepted but unused — nil must be safe.
	if nilGot := buildTopicEntries(nil); len(nilGot) != len(want) {
		t.Errorf("nil fs: expected %d topics, got %d", len(want), len(nilGot))
	}
}
