//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestGroupByFile — ERROR 진단만 파일별 그룹화하고 레이어 분류·순서 보존 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestGroupByFile(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError, File: "db/schema.sql", Message: "e1"},
		{Level: diagnostic.LevelWarning, File: "db/schema.sql", Message: "w1"}, // skipped
		{Level: diagnostic.LevelError, File: "api/openapi.yaml", Message: "e2"},
		{Level: diagnostic.LevelError, File: "db/schema.sql", Message: "e3"},
	}
	groups := groupByFile(diags, "/specs")
	if len(groups) != 2 {
		t.Fatalf("groups = %d, want 2", len(groups))
	}
	// Insertion order preserved: schema.sql first.
	if groups[0].relFile != "db/schema.sql" || groups[0].layer != layerDDL {
		t.Errorf("group0 = %+v, want db/schema.sql/DDL", groups[0])
	}
	if len(groups[0].diags) != 2 {
		t.Errorf("schema.sql diags = %d, want 2 (warning excluded)", len(groups[0].diags))
	}
	if groups[1].relFile != "api/openapi.yaml" || groups[1].layer != layerOpenAPI {
		t.Errorf("group1 = %+v, want api/openapi.yaml/OpenAPI", groups[1])
	}
}
