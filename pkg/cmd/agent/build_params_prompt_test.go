//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestBuildParamsPrompt — feature op/path/desc 가 parameters 프롬프트에 포함되는지 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildParamsPrompt(t *testing.T) {
	feat := features.Feature{Op: "GetWorkflow", Path: "/v1/workflows/{id}", Desc: "fetch one"}
	got := buildParamsPrompt(feat)

	for _, want := range []string{
		"OpenAPI parameters for this endpoint:",
		"op: GetWorkflow",
		"path: /v1/workflows/{id}",
		"desc: fetch one",
		"format: int64",
		"Output ONLY the parameters array",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("expected output to contain %q, got:\n%s", want, got)
		}
	}
}
