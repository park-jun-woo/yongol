//ff:func feature=rule type=test control=sequence
//ff:what populateSSaCSeq — get 분기: model ref + pluralized DDL table 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// TestPopulateSSaCSeq_GetModelRef verifies get/post/put/delete registers model
// ref + pluralized DDL table name.
func TestPopulateSSaCSeq_GetModelRef(t *testing.T) {
	g := newGround()
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)

	seq := ssac.Sequence{Type: "get", Model: "Course.FindByID"}
	populateSSaCSeq(g, "GetCourse", seq, authPairs, callRefs, modelRefs, pubTopics)

	if !modelRefs["Course"] {
		t.Errorf("modelRefs missing Course: %v", modelRefs)
	}
	if !modelRefs["courses"] {
		t.Errorf("modelRefs missing plural 'courses': %v", modelRefs)
	}
}
