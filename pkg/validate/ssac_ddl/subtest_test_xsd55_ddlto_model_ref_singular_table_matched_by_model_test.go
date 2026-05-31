//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXsd55DDLToModelRefSingularTableMatchedByModel — singular table matched by model 서브테스트
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func subtestTestXsd55DDLToModelRefSingularTableMatchedByModel(t *testing.T) {

	fs := newXsd55Fullstack(rule.StringSet{}, ddl.Table{Name: "app_config"})
	fs.ServiceFuncs = []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Model: "AppConfig.Get"}}},
	}
	if diags := xsd55DDLToModelRef(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags for singular table %q, got %d: %v", "app_config", len(diags), diags)
	}

}
