//ff:func feature=agent type=test control=sequence
//ff:what TestBuildSSaCUserPrompt — DDL/쿼리/pathBlock 옵션 섹션 유무에 따른 SSaC 프롬프트 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildSSaCUserPrompt(t *testing.T) {
	feat := features.Feature{Op: "CreateWorkflow", Path: "/v1/workflows", Table: "workflows", Public: false, Desc: "make one"}

	t.Run("AllSections", func(t *testing.T) {
		got := buildSSaCUserPrompt(feat, "CREATE TABLE workflows ();", []string{"CreateWorkflow", "GetWorkflow"}, "/v1/workflows:\n  post:")
		for _, want := range []string{
			"op: CreateWorkflow",
			"public: false",
			"DDL for table workflows:\nCREATE TABLE workflows",
			"Available sqlc queries:",
			"  CreateWorkflow",
			"  GetWorkflow",
			"OpenAPI path block:\n/v1/workflows:",
			"Generate a single SSaC service file",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("NoOptionalSections", func(t *testing.T) {
		got := buildSSaCUserPrompt(feat, "", nil, "")
		for _, unwanted := range []string{"DDL for table", "Available sqlc queries:", "OpenAPI path block:"} {
			if strings.Contains(got, unwanted) {
				t.Errorf("did not expect %q, got:\n%s", unwanted, got)
			}
		}
		if !strings.Contains(got, "op: CreateWorkflow") {
			t.Errorf("expected feature section, got:\n%s", got)
		}
	})
}
