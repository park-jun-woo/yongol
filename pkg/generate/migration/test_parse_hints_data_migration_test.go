//ff:func feature=migration type=test control=sequence
//ff:what TestParseHints_DataMigration — @data_migration file=... 힌트가 DataMigrations 맵에 등록
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestParseHints_DataMigration(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "data_migration", Args: map[string]string{"file": "migrations_data/0042.sql"}, TableCtx: "users", BlockAbove: true},
	}
	h := ParseHints(comments)
	if got := h.DataMigrations["users"]; got != "migrations_data/0042.sql" {
		t.Errorf("data_migration wrong: %q", got)
	}
}
