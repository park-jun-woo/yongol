//ff:func feature=rule type=test control=sequence
//ff:what populateSSaCSeq — auth 분기: action:resource 가 authPairs 에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// TestPopulateSSaCSeq_Auth verifies auth action:resource is appended.
func TestPopulateSSaCSeq_Auth(t *testing.T) {
	g := newGround()
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)

	seq := ssac.Sequence{Type: "auth", Action: "delete", Resource: "project"}
	populateSSaCSeq(g, "DeleteProject", seq, authPairs, callRefs, modelRefs, pubTopics)

	if !authPairs["delete:project"] {
		t.Errorf("authPairs missing delete:project: %v", authPairs)
	}
}
