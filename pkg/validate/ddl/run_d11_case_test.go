//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ddl-structural
//ff:what runD11Case — TestD11UnsupportedPgType 의 단일 케이스 실행 (Fullstack 빌드 + 진단 비교)

package ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runD11Case(t *testing.T, columns map[string]ddl.Column, want int) {
	t.Helper()
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{{Name: "t", File: "t.sql", Line: 1, Columns: columns}},
	}
	diags := d11UnsupportedPgType(fs)
	if len(diags) != want {
		t.Errorf("len(diags) = %d, want %d", len(diags), want)
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, "[D-11]") {
			t.Errorf("diag message %q missing [D-11] prefix", d.Message)
		}
	}
}
