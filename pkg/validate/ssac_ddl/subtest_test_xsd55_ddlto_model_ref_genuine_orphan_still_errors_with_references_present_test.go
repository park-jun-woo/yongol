//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXsd55DDLToModelRefGenuineOrphanStillErrorsWithReferencesPresent — genuine orphan still errors with references present 서브테스트
package ssac_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func subtestTestXsd55DDLToModelRefGenuineOrphanStillErrorsWithReferencesPresent(t *testing.T) {

	fs := newXsd55Fullstack(
		rule.StringSet{},
		ddl.Table{Name: "users"},
		ddl.Table{Name: "abandoned_widgets"},
	)
	fs.ServiceFuncs = []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Model: "User.Get"}}},
	}
	diags := xsd55DDLToModelRef(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "abandoned_widgets") {
		t.Errorf("expected abandoned_widgets to error, got: %q", diags[0].Message)
	}

}
