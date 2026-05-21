//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what XFD-02 — belongs_to FK 컬럼이 없을 때 ERROR 진단 테스트
package features_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFD02_FKExists_Fires(t *testing.T) {
	ft := map[string]features.TableDef{
		"tasks": {BelongsTo: []string{"workflows"}},
	}
	tables := []ddl.Table{
		{Name: "tasks", Columns: map[string]ddl.Column{
			"id":   {Name: "id", RawType: "BIGINT"},
			"name": {Name: "name", RawType: "TEXT"},
		}},
	}
	fs := buildFSForXFD(ft, tables)
	diags := xfd02FKExists(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XFD-02]") {
		t.Errorf("want [XFD-02] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "workflows_id") {
		t.Errorf("want FK column name in message, got %s", diags[0].Message)
	}
}
