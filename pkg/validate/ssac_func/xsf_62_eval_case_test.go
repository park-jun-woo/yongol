//ff:type feature=validate type=model
//ff:what xsf62EvalCase — TestXsf62EvalOnlyRef table-driven 케이스 정의

package ssac_func

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// xsf62EvalCase captures one row of the BUG-002 XSF-62 integration table.
// seqs populate a single ServiceFunc; wantDiags is the expected diagnostic
// count after running xsf62FuncSpecUsed against a Func Spec billing.isZeroBalance.
type xsf62EvalCase struct {
	name      string
	seqs      []ssac.Sequence
	wantDiags int
}
