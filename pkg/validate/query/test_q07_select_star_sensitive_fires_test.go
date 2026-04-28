//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what Q-07 positive — @sensitive 컬럼이 있는 테이블에 SELECT * 면 WARNING

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ07SelectStarSensitiveFires covers Q-07: SELECT * on a table with a
// @sensitive column emits a WARNING.
func TestQ07SelectStarSensitiveFires(t *testing.T) {
	specs := loadSpecs(t, "q07_select_star.sql")
	fs := &yongol.Fullstack{
		SQLcQueries: specs,
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"password_hash": {Name: "password_hash", RawType: "TEXT", Sensitive: true},
			},
		}},
	}
	diags := q07SelectStarSensitive(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-07 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-07]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}
