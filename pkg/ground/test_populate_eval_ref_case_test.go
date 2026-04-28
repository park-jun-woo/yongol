//ff:type feature=rule type=model
//ff:what populateEvalRefCase — TestPopulateSSaCEvalRef table-driven 케이스 정의

package ground

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// populateEvalRefCase captures one row of the BUG-002 populate-eval table.
// seqs are fed sequentially through populateSSaCSeq; wantRefs are asserted
// against the resulting callRefs StringSet.
type populateEvalRefCase struct {
	name     string
	seqs     []ssac.Sequence
	wantRefs []string
}
