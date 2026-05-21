//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what XFD-01 — features table이 DDL에 없을 때 ERROR 진단 테스트
package features_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFD01_TableExists_Fires(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {},
		"tasks":     {},
	}
	tables := []ddl.Table{
		{Name: "workflows"},
	}
	fs := buildFSForXFD(ft, tables)
	diags := xfd01TableExists(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XFD-01]") {
		t.Errorf("want [XFD-01] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "tasks") {
		t.Errorf("want table name in message, got %s", diags[0].Message)
	}
}
