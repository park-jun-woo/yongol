//ff:func feature=agent type=test control=sequence
//ff:what TestBuildRetryPrompt — DDL/relativeLine 유무에 따른 재시도 프롬프트 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildRetryPrompt(t *testing.T) {
	feat := features.Feature{Op: "CreateWorkflow", Path: "/v1/workflows", Table: "workflows", Public: true, Desc: "make one"}

	t.Run("WithDDLAndLine", func(t *testing.T) {
		got := buildRetryPrompt(feat, "CREATE TABLE workflows ();", "boom", 12)
		for _, want := range []string{
			"op: CreateWorkflow",
			"public: true",
			"Table DDL:\nCREATE TABLE workflows",
			"Previous attempt had this error:\nboom",
			"near line 12",
			"corrected OpenAPI path block",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("NoDDLNoLine", func(t *testing.T) {
		got := buildRetryPrompt(feat, "", "oops", -1)
		if strings.Contains(got, "Table DDL:") {
			t.Errorf("did not expect DDL section, got:\n%s", got)
		}
		if strings.Contains(got, "near line") {
			t.Errorf("did not expect line hint for negative relativeLine, got:\n%s", got)
		}
		if !strings.Contains(got, "Previous attempt had this error:\noops") {
			t.Errorf("expected prev error, got:\n%s", got)
		}
	})
}
