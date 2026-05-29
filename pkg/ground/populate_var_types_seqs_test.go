//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateVarTypesSeqs — Result.Var → 원본 Type spec 매핑 (wrapper/prefix 보존)

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestPopulateVarTypesSeqs_PreserveRawType verifies the original type string
// (including "[]", "*", "pkg." prefix) is stored verbatim under SSaC.var.*.
func TestPopulateVarTypesSeqs_PreserveRawType(t *testing.T) {
	seqs := []ssac.Sequence{
		{Result: &ssac.Result{Var: "course", Type: "Course"}},
		{Result: &ssac.Result{Var: "webhooks", Type: "[]Webhook"}},
		{Result: &ssac.Result{Var: "u", Type: "*User"}},
		{Result: &ssac.Result{Var: "resp", Type: "billing.CheckCreditsResponse"}},
		// Missing var → must be skipped
		{Result: &ssac.Result{Var: "", Type: "SomeType"}},
		// nil Result → must be skipped
		{Result: nil},
	}
	g := newGround()
	populateVarTypesSeqs(g, "GetCourse", seqs)

	if g.Types["SSaC.var.GetCourse.course"] != "Course" {
		t.Errorf("course type lost")
	}
	if g.Types["SSaC.var.GetCourse.webhooks"] != "[]Webhook" {
		t.Errorf("webhooks type lost slice prefix: %q", g.Types["SSaC.var.GetCourse.webhooks"])
	}
	if g.Types["SSaC.var.GetCourse.u"] != "*User" {
		t.Errorf("u type lost pointer prefix: %q", g.Types["SSaC.var.GetCourse.u"])
	}
	if g.Types["SSaC.var.GetCourse.resp"] != "billing.CheckCreditsResponse" {
		t.Errorf("resp type lost package prefix: %q", g.Types["SSaC.var.GetCourse.resp"])
	}
}
