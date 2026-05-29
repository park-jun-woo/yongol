//ff:func feature=rule type=test control=sequence
//ff:what collectQueryFields — args 에서 query 필드만 dst 에 수집

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestCollectQueryFields_FromArgs(t *testing.T) {
	args := []parseArg{
		{Source: "query", Field: "page"},
		{Source: "query", Field: "size"},
		{Source: "request", Field: "title"}, // not query, must be skipped
	}
	dst := make(rule.StringSet)
	collectQueryFields(args, nil, dst)
	if !dst["page"] || !dst["size"] {
		t.Errorf("query fields missing: %v", dst)
	}
	if dst["title"] {
		t.Errorf("non-query field should not be included: %v", dst)
	}
}
