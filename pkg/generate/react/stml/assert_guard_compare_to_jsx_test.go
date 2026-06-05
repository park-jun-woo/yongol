//ff:func feature=stml-gen type=test-helper control=sequence
//ff:what guardCompareToJSX 단일 케이스 변환 결과 검증 헬퍼

package stml

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// assertGuardCompareToJSX builds a compare GuardExpr from ref/op/value and
// asserts guardCompareToJSX renders the expected JSX string.
func assertGuardCompareToJSX(t *testing.T, ref stml.GuardRef, op, value, dataVar, want string) {
	t.Helper()
	expr := &stml.GuardExpr{
		Kind:  stml.GuardCompare,
		Ref:   ref,
		Op:    op,
		Value: value,
	}
	got := guardCompareToJSX(expr, dataVar)
	if got != want {
		t.Errorf("guardCompareToJSX()\n  got:  %s\n  want: %s", got, want)
	}
}
