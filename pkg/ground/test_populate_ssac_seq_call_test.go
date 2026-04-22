//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateSSaCSeq — @auth/@call/@publish/@get 분기별 레지스트리 등록

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

// TestPopulateSSaCSeq_Publish verifies publish topic is registered.
func TestPopulateSSaCSeq_Publish(t *testing.T) {
	g := newGround()
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)

	seq := ssac.Sequence{Type: "publish", Topic: "order.completed"}
	populateSSaCSeq(g, "CheckoutOrder", seq, authPairs, callRefs, modelRefs, pubTopics)

	if !pubTopics["order.completed"] {
		t.Errorf("pubTopics missing order.completed: %v", pubTopics)
	}
}

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
