//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what XFD-02 — DDL 테이블 누락 시 스킵 테스트 (XFD-01이 잡으므로)
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFD02_FKExists_MissingDDL_Skips(t *testing.T) {
	ft := map[string]features.TableDef{
		"tasks": {BelongsTo: []string{"workflows"}},
	}
	// DDL has no "tasks" table — XFD-01 will catch it; XFD-02 should skip.
	tables := []ddl.Table{
		{Name: "workflows"},
	}
	fs := buildFSForXFD(ft, tables)
	diags := xfd02FKExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags when DDL table is missing, got %d", len(diags))
	}
}
