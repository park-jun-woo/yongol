//ff:func feature=rule type=test control=sequence
//ff:what populateSSaCSeq — response 분기: @response fields 가 Schemas에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// TestPopulateSSaCSeq_ResponseFields verifies @response type registers fields
// via populateResponseFields.
func TestPopulateSSaCSeq_ResponseFields(t *testing.T) {
	g := newGround()
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)

	seq := ssac.Sequence{
		Type:   "response",
		Fields: map[string]string{"course": "course", "name": "course.Name"},
	}
	populateSSaCSeq(g, "GetCourse", seq, authPairs, callRefs, modelRefs, pubTopics)

	got := g.Schemas["SSaC.response.GetCourse"]
	if len(got) != 2 {
		t.Errorf("SSaC.response.GetCourse = %v, want 2 fields", got)
	}
}
