//ff:func feature=rule type=test control=sequence
//ff:what populateSSaCSeq — publish 분기: topic 이 pubTopics 에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

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
