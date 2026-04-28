//ff:func feature=validate type=test-helper control=sequence topic=func-check
//ff:what runXsf62EvalCase — 단일 xsf62EvalCase 실행 + 진단 개수 검증

package ssac_func

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
)

// runXsf62EvalCase builds a Fullstack with one ServiceFunc carrying tc.seqs
// plus a Func Spec billing.isZeroBalance, runs ground.Build, then asserts
// xsf62FuncSpecUsed produces exactly tc.wantDiags diagnostics.
func runXsf62EvalCase(t *testing.T, tc xsf62EvalCase) {
	t.Helper()
	fs := buildXsf62EvalFullstack(tc.seqs)
	fs.SetGround(ground.Build(fs))
	diags := xsf62FuncSpecUsed(fs)
	if len(diags) != tc.wantDiags {
		t.Fatalf("%s: want %d diags, got %d (%+v)", tc.name, tc.wantDiags, len(diags), diags)
	}
	if tc.wantDiags == 1 && !strings.Contains(diags[0].Message, "[XSF-62]") {
		t.Errorf("%s: rule id missing in message: %q", tc.name, diags[0].Message)
	}
}
