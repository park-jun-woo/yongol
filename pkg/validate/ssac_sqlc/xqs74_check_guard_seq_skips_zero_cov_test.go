//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS74Branches_ZeroCov — xqs74CheckFunc/GuardSeq/Model 잔여 분기 직접 커버
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXQS74CheckGuardSeq_Skips_ZeroCov(t *testing.T) {
	varModel := map[string]string{"u": "User"}
	tableMap := map[string]*ddl.Table{}

	// non-guard sequence type
	if _, ok := xqs74CheckGuardSeq(ssacparser.Sequence{Type: "get", Target: "u"}, varModel, tableMap, "f"); ok {
		t.Error("get type should be skipped")
	}
	// empty target
	if _, ok := xqs74CheckGuardSeq(ssacparser.Sequence{Type: "empty", Target: ""}, varModel, tableMap, "f"); ok {
		t.Error("empty target should be skipped")
	}
	// dotted target
	if _, ok := xqs74CheckGuardSeq(ssacparser.Sequence{Type: "exists", Target: "u.ID"}, varModel, tableMap, "f"); ok {
		t.Error("dotted target should be skipped")
	}
	// unknown var
	if _, ok := xqs74CheckGuardSeq(ssacparser.Sequence{Type: "empty", Target: "missing"}, varModel, tableMap, "f"); ok {
		t.Error("unknown var should be skipped")
	}
}
