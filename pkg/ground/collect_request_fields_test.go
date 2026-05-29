//ff:func feature=rule type=test control=iteration dimension=1
//ff:what collectRequestFields — args+inputs 의 request.* 참조를 dst 에 수집

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

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
