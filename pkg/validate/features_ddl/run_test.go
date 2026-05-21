//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what Run — Features↔DDL 교차 검증 두 규칙 동시 발동 테스트
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestRun_BothRulesFire(t *testing.T) {
	// "missing_table" is in FeatureTables but not in DDL (XFD-01).
	// "tasks" belongs_to "workflows" but DDL has no "workflows_id" column (XFD-02).
	ft := map[string]features.TableDef{
		"missing_table": {},
		"tasks":         {BelongsTo: []string{"workflows"}},
	}
	tables := []ddl.Table{
		{Name: "tasks", Columns: map[string]ddl.Column{
			"id": {Name: "id", RawType: "BIGINT"},
		}},
	}
	fs := buildFSForXFD(ft, tables)
	diags := Run(fs)
	if len(diags) != 2 {
		t.Fatalf("want 2 diags (one XFD-01, one XFD-02), got %d", len(diags))
	}
}
