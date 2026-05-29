//ff:func feature=rule type=test control=iteration dimension=1
//ff:what populateSSaCSeq — eval 분기: SSaC @eval 의 Func 참조가 callRefs 에 등록 (BUG-002)

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestPopulateSSaCEvalRef verifies BUG-002 fix: @eval registers the referenced
// Func into SSaC.callRef (PascalCase → camelCase normalization, identical to
// @call). Cases:
//  1. @eval-only sequence       → callRefs contains the eval target
//  2. @call + @eval (different) → callRefs contains both
//  3. @call-only (regression)   → callRefs contains the call target
func TestPopulateSSaCEvalRef(t *testing.T) {
	cases := []populateEvalRefCase{
		{
			name:     "eval-only registers callRef",
			seqs:     []ssac.Sequence{{Type: ssac.SeqEval, Model: "billing.IsZeroBalance"}},
			wantRefs: []string{"billing.isZeroBalance"},
		},
		{
			name: "call and eval both register",
			seqs: []ssac.Sequence{
				{Type: ssac.SeqCall, Model: "auth.HashPassword"},
				{Type: ssac.SeqEval, Model: "billing.IsZeroBalance"},
			},
			wantRefs: []string{"auth.hashPassword", "billing.isZeroBalance"},
		},
		{
			name:     "call-only regression",
			seqs:     []ssac.Sequence{{Type: ssac.SeqCall, Model: "auth.HashPassword"}},
			wantRefs: []string{"auth.hashPassword"},
		},
	}
	for _, tc := range cases {
		runPopulateEvalRefCase(t, tc)
	}
}
