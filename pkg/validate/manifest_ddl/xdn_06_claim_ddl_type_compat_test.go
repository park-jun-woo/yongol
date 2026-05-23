//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what claimDDLTypeCompatible — 알 수 없는 claim 타입 + 매트릭스 매칭/불매칭 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestClaimDDLTypeCompatible(t *testing.T) {
	tests := []struct {
		name      string
		claimType string
		col       ddl.Column
		want      bool
	}{
		{
			name:      "unknown claim type returns false",
			claimType: "float64",
			col:       ddl.Column{RawType: "NUMERIC"},
			want:      false,
		},
		{
			name:      "int64 matches BIGINT",
			claimType: "int64",
			col:       ddl.Column{RawType: "BIGINT"},
			want:      true,
		},
		{
			name:      "int64 matches INT8",
			claimType: "int64",
			col:       ddl.Column{RawType: "INT8"},
			want:      true,
		},
		{
			name:      "int64 does not match TEXT",
			claimType: "int64",
			col:       ddl.Column{RawType: "TEXT"},
			want:      false,
		},
		{
			name:      "string matches TEXT",
			claimType: "string",
			col:       ddl.Column{RawType: "TEXT"},
			want:      true,
		},
		{
			name:      "string matches VARCHAR(255)",
			claimType: "string",
			col:       ddl.Column{RawType: "VARCHAR(255)"},
			want:      true,
		},
		{
			name:      "bool matches BOOLEAN",
			claimType: "bool",
			col:       ddl.Column{RawType: "BOOLEAN"},
			want:      true,
		},
		{
			name:      "uuid matches UUID",
			claimType: "uuid",
			col:       ddl.Column{RawType: "UUID"},
			want:      true,
		},
		{
			name:      "uuid does not match TEXT",
			claimType: "uuid",
			col:       ddl.Column{RawType: "TEXT"},
			want:      false,
		},
		{
			name:      "int32 matches INTEGER",
			claimType: "int32",
			col:       ddl.Column{RawType: "INTEGER"},
			want:      true,
		},
		{
			name:      "int32 does not match BIGINT",
			claimType: "int32",
			col:       ddl.Column{RawType: "BIGINT"},
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := claimDDLTypeCompatible(tt.claimType, tt.col)
			if got != tt.want {
				t.Errorf("claimDDLTypeCompatible(%q, %q) = %v, want %v",
					tt.claimType, tt.col.RawType, got, tt.want)
			}
		})
	}
}
