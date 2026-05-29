//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what Run nil 가드 — DesignSpec nil 시 빈 결과 반환

package design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunNilGuard(t *testing.T) {
	fs := &yongol.Fullstack{DesignSpec: nil}
	if got := Run(fs); len(got) != 0 {
		t.Fatalf("expected nil result, got %d diagnostics", len(got))
	}
}
