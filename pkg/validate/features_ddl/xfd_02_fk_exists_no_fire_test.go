//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what XFD-02 — FK 컬럼이 있을 때 정상 통과 테스트
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFD02_FKExists_NoFire(t *testing.T) {
	ft := map[string]features.TableDef{
		"tasks": {BelongsTo: []string{"workflows"}},
	}
	tables := []ddl.Table{
		{Name: "tasks", Columns: map[string]ddl.Column{
			"id":           {Name: "id", RawType: "BIGINT"},
			"workflows_id": {Name: "workflows_id", RawType: "BIGINT"},
		}},
	}
	fs := buildFSForXFD(ft, tables)
	diags := xfd02FKExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
