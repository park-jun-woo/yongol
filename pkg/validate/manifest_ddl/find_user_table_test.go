//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what findUserTable — nil/empty 조기 반환 + 이름 일치/불일치 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFindUserTable(t *testing.T) {
	tables := []ddl.Table{
		{Name: "users"},
		{Name: "orders"},
		{Name: "products"},
	}

	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		userTable string
		wantName  string
		wantNil   bool
	}{
		{
			name:      "nil fullstack returns nil",
			fs:        nil,
			userTable: "users",
			wantNil:   true,
		},
		{
			name:      "empty userTable returns nil",
			fs:        &yongol.Fullstack{DDLTables: tables},
			userTable: "",
			wantNil:   true,
		},
		{
			name:      "matching table found",
			fs:        &yongol.Fullstack{DDLTables: tables},
			userTable: "users",
			wantName:  "users",
			wantNil:   false,
		},
		{
			name:      "matching table in middle",
			fs:        &yongol.Fullstack{DDLTables: tables},
			userTable: "orders",
			wantName:  "orders",
			wantNil:   false,
		},
		{
			name:      "no matching table returns nil",
			fs:        &yongol.Fullstack{DDLTables: tables},
			userTable: "nonexistent",
			wantNil:   true,
		},
		{
			name:      "empty DDLTables returns nil",
			fs:        &yongol.Fullstack{},
			userTable: "users",
			wantNil:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFindUserTable(t, tt.fs, tt.userTable, tt.wantNil, tt.wantName)
		})
	}
}
