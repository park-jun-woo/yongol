//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildXqs18DDLColumnTypeMap_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"id":   {Name: "id", RawType: "BIGINT", NotNull: true},
				"name": {Name: "name", RawType: "TEXT"},
			},
		}},
	}
	m := buildXqs18DDLColumnTypeMap(fs)
	cols := m["users"]
	if cols == nil {
		t.Fatalf("expected users table, got %v", m)
	}
	if cols["id"] == "" || cols["name"] == "" {
		t.Errorf("expected Go types for columns, got %v", cols)
	}
}
