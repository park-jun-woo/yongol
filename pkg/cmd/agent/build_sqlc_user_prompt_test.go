//ff:func feature=agent type=test control=selection dimension=1
//ff:what TestBuildSQLcUserPrompt — 관련 feature 유무에 따른 sqlc 쿼리 프롬프트 구성 검증

package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildSQLcUserPrompt(t *testing.T) {
	t.Run("WithFeatures", func(t *testing.T) {
		feats := []features.Feature{
			{Op: "GetWorkflow", Path: "/v1/workflows/{id}", Desc: "fetch"},
		}
		got := buildSQLcUserPrompt("workflows", "CREATE TABLE workflows ();", feats)
		for _, want := range []string{
			"Table: workflows",
			"DDL:\nCREATE TABLE workflows",
			"Related features with cardinality hints:",
			"GetWorkflow /v1/workflows/{id}: fetch (cardinality:",
			"Generate sqlc-compatible SQL queries",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("WithoutFeatures", func(t *testing.T) {
		got := buildSQLcUserPrompt("plain", "CREATE TABLE plain ();", nil)
		if strings.Contains(got, "Related features") {
			t.Errorf("did not expect features section, got:\n%s", got)
		}
		if !strings.Contains(got, "Table: plain") {
			t.Errorf("expected table line, got:\n%s", got)
		}
	})
}
