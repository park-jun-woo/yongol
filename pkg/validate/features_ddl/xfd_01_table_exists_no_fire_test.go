//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what XFD-01 — features table이 DDL에 모두 있을 때 정상 통과 테스트
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFD01_TableExists_NoFire(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {},
		"tasks":     {},
	}
	tables := []ddl.Table{
		{Name: "workflows"},
		{Name: "tasks"},
	}
	fs := buildFSForXFD(ft, tables)
	diags := xfd01TableExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
