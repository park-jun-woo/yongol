//ff:func feature=stml-gen type=test-helper control=sequence
//ff:what guardLifecycleToJSX 단일 케이스 변환 결과 검증 헬퍼

package stml

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// assertGuardLifecycleToJSX builds a lifecycle GuardExpr from ref/lifecycle and
// asserts guardLifecycleToJSX renders the expected JSX string.
func assertGuardLifecycleToJSX(t *testing.T, ref stml.GuardRef, lifecycle, dataVar, want string) {
	t.Helper()
	expr := &stml.GuardExpr{
		Kind:      stml.GuardLifecycle,
		Ref:       ref,
		Lifecycle: lifecycle,
	}
	got := guardLifecycleToJSX(expr, dataVar)
	if got != want {
		t.Errorf("guardLifecycleToJSX()\n  got:  %s\n  want: %s", got, want)
	}
}
