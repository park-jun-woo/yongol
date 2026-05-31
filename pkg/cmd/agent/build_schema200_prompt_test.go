//ff:func feature=agent type=test control=sequence
//ff:what TestBuildSchema200Prompt — DDL 컨텍스트 유무에 따른 200 schema 프롬프트 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildSchema200Prompt(t *testing.T) {
	feat := features.Feature{Op: "GetWorkflow", Desc: "fetch one"}

	t.Run("WithDDL", func(t *testing.T) {
		got := buildSchema200Prompt(feat, "CREATE TABLE workflows (id BIGINT);")
		for _, want := range []string{
			"op: GetWorkflow",
			"desc: fetch one",
			"DDL:\nCREATE TABLE workflows",
			"format: int64",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("WithoutDDL", func(t *testing.T) {
		got := buildSchema200Prompt(feat, "")
		if strings.Contains(got, "DDL:") {
			t.Errorf("did not expect DDL section, got:\n%s", got)
		}
		if !strings.Contains(got, "Output ONLY the schema") {
			t.Errorf("expected instruction, got:\n%s", got)
		}
	})
}
