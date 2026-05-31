//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXsd55DDLToModelRefArchivedExempt — archived exempt 서브테스트
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func subtestTestXsd55DDLToModelRefArchivedExempt(t *testing.T) {

	fs := newXsd55Fullstack(
		rule.StringSet{"archived.legacy": true},
		ddl.Table{Name: "legacy"},
	)
	if diags := xsd55DDLToModelRef(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags for archived table, got %d: %v", len(diags), diags)
	}

}
