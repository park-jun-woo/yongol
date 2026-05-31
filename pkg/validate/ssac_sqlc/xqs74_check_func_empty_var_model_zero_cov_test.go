//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS74Branches_ZeroCov — xqs74CheckFunc/GuardSeq/Model 잔여 분기 직접 커버
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXQS74CheckFunc_EmptyVarModel_ZeroCov(t *testing.T) {
	// Function with only guard sequences and no get/post result bindings →
	// varModel is empty → early nil return.
	fn := ssacparser.ServiceFunc{
		Name:     "GuardOnly",
		FileName: "guard.ssac",
		Sequences: []ssacparser.Sequence{
			{Type: "empty", Target: "x", Line: 2},
			{Type: "get", Model: "A.B"}, // get with nil Result → skipped, no binding
		},
	}
	if diags := xqs74CheckFunc(fn, map[string]*ddl.Table{}); diags != nil {
		t.Fatalf("expected nil diags for empty varModel, got %v", diags)
	}
}
