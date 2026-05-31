//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXsd55DDLToModelRefOrphanTableErrors — orphan table errors 서브테스트
package ssac_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func subtestTestXsd55DDLToModelRefOrphanTableErrors(t *testing.T) {

	fs := newXsd55Fullstack(rule.StringSet{}, ddl.Table{Name: "orphans"})
	diags := xsd55DDLToModelRef(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "orphans") {
		t.Errorf("diag message missing table name: %q", diags[0].Message)
	}

}
