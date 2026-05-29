//ff:func feature=rule type=test control=sequence
//ff:what collectQueryFields — inputs 의 "query.*" 참조만 dst 에 수집

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestCollectQueryFields_FromInputs(t *testing.T) {
	inputs := map[string]string{
		"cursor": "query.cursor",
		"body":   "request.body",
	}
	dst := make(rule.StringSet)
	collectQueryFields(nil, inputs, dst)
	if !dst["cursor"] {
		t.Errorf("cursor not stripped from 'query.cursor': %v", dst)
	}
	if dst["body"] {
		t.Errorf("request.* should not leak into query set: %v", dst)
	}
}
