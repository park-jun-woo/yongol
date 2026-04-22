//ff:func feature=rule type=test control=sequence dimension=1
//ff:what collectQueryFields / collectRequestFields — query.*/request.* 참조 집합 추출

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

func TestCollectRequestFields_FromArgsAndInputs(t *testing.T) {
	args := []parseArg{
		{Source: "request", Field: "title"},
		{Source: "request", Field: "body"},
	}
	inputs := map[string]string{
		"authorId": "request.author_id",
		"pageKey":  "query.page", // must be skipped
	}
	dst := make(rule.StringSet)
	collectRequestFields(args, inputs, dst)

	for _, f := range []string{"title", "body", "author_id"} {
		if !dst[f] {
			t.Errorf("request fields missing %q: %v", f, dst)
		}
	}
	if dst["page"] {
		t.Errorf("query field must not leak into request set: %v", dst)
	}
}
