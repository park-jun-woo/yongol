//ff:func feature=rule type=test control=sequence
//ff:what populateSSaCSeq — call 분기: "pkg.Name" → "pkg.name" 로 정규화

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// TestPopulateSSaCSeq_Call verifies call "pkg.Name" is normalized to
// "pkg.name" and inserted into callRefs.
func TestPopulateSSaCSeq_Call(t *testing.T) {
	g := newGround()
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)

	seq := ssac.Sequence{Type: "call", Model: "auth.HashPassword"}
	populateSSaCSeq(g, "Signup", seq, authPairs, callRefs, modelRefs, pubTopics)

	if !callRefs["auth.hashPassword"] {
		t.Errorf("callRefs missing auth.hashPassword (PascalCase→camelCase): %v", callRefs)
	}
}
