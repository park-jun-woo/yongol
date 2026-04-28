//ff:func feature=rule type=test-helper control=iteration dimension=1
//ff:what runPopulateEvalRefCase — 단일 populateEvalRefCase 실행 + callRefs 검증

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// runPopulateEvalRefCase executes one populateEvalRefCase: feeds sequences
// through populateSSaCSeq and asserts that every wantRefs entry appears in
// the resulting callRefs StringSet.
func runPopulateEvalRefCase(t *testing.T, tc populateEvalRefCase) {
	t.Helper()
	g := newGround()
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)
	for _, s := range tc.seqs {
		populateSSaCSeq(g, "Charge", s, authPairs, callRefs, modelRefs, pubTopics)
	}
	for _, want := range tc.wantRefs {
		if !callRefs[want] {
			t.Errorf("%s: callRefs missing %q: %v", tc.name, want, callRefs)
		}
	}
}
