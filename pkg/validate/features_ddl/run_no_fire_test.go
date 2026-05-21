//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what Run — Features↔DDL 교차 검증 정상 통과 테스트
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestRun_NoFire(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {},
		"tasks":     {BelongsTo: []string{"workflows"}},
	}
	tables := []ddl.Table{
		{Name: "workflows"},
		{Name: "tasks", Columns: map[string]ddl.Column{
			"id":           {Name: "id", RawType: "BIGINT"},
			"workflows_id": {Name: "workflows_id", RawType: "BIGINT"},
		}},
	}
	fs := buildFSForXFD(ft, tables)
	diags := Run(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags when all tables and FKs match, got %d", len(diags))
	}
}
