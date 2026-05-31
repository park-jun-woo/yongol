//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestResolveQueryName(t *testing.T) {
	if got := resolveQueryName(ssacparser.Sequence{Model: "Workflow.FindByID"}); got != "WorkflowFindByID" {
		t.Errorf("got %q, want WorkflowFindByID", got)
	}
	if got := resolveQueryName(ssacparser.Sequence{Model: "Get"}); got != "Get" {
		t.Errorf("got %q, want Get", got)
	}
}
