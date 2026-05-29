//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what XSF-62 — @eval 만으로 참조되는 Func Spec 은 미사용으로 잘못 판정되지 않아야 함 (BUG-002)

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestXsf62EvalOnlyRef drives ground.Build with real ServiceFuncs / Sequences
// to verify BUG-002: a Func Spec referenced exclusively via @eval must not be
// flagged as unused. Three cases:
//  1. @eval-only reference   → 0 diagnostics
//  2. no reference at all    → 1 diagnostic (regression guard)
//  3. @call-only reference   → 0 diagnostics (regression guard)
func TestXsf62EvalOnlyRef(t *testing.T) {
	cases := []xsf62EvalCase{
		{
			name:      "eval-only reference passes",
			seqs:      []ssac.Sequence{{Type: ssac.SeqEval, Model: "billing.IsZeroBalance"}},
			wantDiags: 0,
		},
		{
			name:      "no reference fires diagnostic",
			seqs:      nil,
			wantDiags: 1,
		},
		{
			name:      "call-only reference passes (regression)",
			seqs:      []ssac.Sequence{{Type: ssac.SeqCall, Model: "billing.IsZeroBalance"}},
			wantDiags: 0,
		},
	}
	for _, tc := range cases {
		runXsf62EvalCase(t, tc)
	}
}
