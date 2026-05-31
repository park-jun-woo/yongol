//ff:func feature=agent type=test control=sequence
//ff:what TestBuildRequestBodyPrompt — DDL 컨텍스트 유무에 따른 requestBody 프롬프트 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildRequestBodyPrompt(t *testing.T) {
	feat := features.Feature{Op: "CreateWorkflow", Path: "/v1/workflows", Desc: "make one"}

	t.Run("WithDDL", func(t *testing.T) {
		got := buildRequestBodyPrompt(feat, "CREATE TABLE workflows (id BIGINT);")
		for _, want := range []string{
			"op: CreateWorkflow",
			"path: /v1/workflows",
			"desc: make one",
			"DDL:\nCREATE TABLE workflows",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("WithoutDDL", func(t *testing.T) {
		got := buildRequestBodyPrompt(feat, "")
		if strings.Contains(got, "DDL:") {
			t.Errorf("did not expect DDL section when ddlContent empty, got:\n%s", got)
		}
		if !strings.Contains(got, "Output \"none\"") {
			t.Errorf("expected instruction, got:\n%s", got)
		}
	})
}
