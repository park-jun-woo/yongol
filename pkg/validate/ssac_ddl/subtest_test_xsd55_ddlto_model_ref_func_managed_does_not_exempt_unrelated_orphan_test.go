//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXsd55DDLToModelRefFuncManagedDoesNotExemptUnrelatedOrphan — func-managed does not exempt unrelated orphan 서브테스트
package ssac_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func subtestTestXsd55DDLToModelRefFuncManagedDoesNotExemptUnrelatedOrphan(t *testing.T) {

	fs := newXsd55Fullstack(
		rule.StringSet{"funcManaged.bids": true},
		ddl.Table{Name: "bids"},
		ddl.Table{Name: "orphans"},
	)
	diags := xsd55DDLToModelRef(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "orphans") {
		t.Errorf("expected orphans to error, got: %q", diags[0].Message)
	}

}
