//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what modelForSQLCMethod 단위 테스트 (sqlc query name → Model, 미스매치 시 "")

package ssac

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

func TestModelForSQLCMethod(t *testing.T) {
	g := &methodGen{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "WorkflowFindByID", Model: "Workflow"},
			{Name: "UserCreate", Model: "User"},
		},
	}
	cases := map[string]string{
		"WorkflowFindByID": "Workflow",
		"UserCreate":       "User",
		"Unknown":          "",
	}
	for method, want := range cases {
		if got := g.modelForSQLCMethod(method); got != want {
			t.Errorf("modelForSQLCMethod(%q) = %q, want %q", method, got, want)
		}
	}
}
