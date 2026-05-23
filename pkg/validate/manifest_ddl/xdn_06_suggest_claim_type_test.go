//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what suggestClaimType — 매트릭스 매칭 + 알 수 없는 타입 기본값 string 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestSuggestClaimType(t *testing.T) {
	tests := []struct {
		name string
		col  ddl.Column
		want string
	}{
		{
			name: "BIGINT suggests int64",
			col:  ddl.Column{RawType: "BIGINT"},
			want: "int64",
		},
		{
			name: "INTEGER suggests int32",
			col:  ddl.Column{RawType: "INTEGER"},
			want: "int32",
		},
		{
			name: "TEXT suggests string",
			col:  ddl.Column{RawType: "TEXT"},
			want: "string",
		},
		{
			name: "UUID suggests uuid",
			col:  ddl.Column{RawType: "UUID"},
			want: "uuid",
		},
		{
			name: "BOOLEAN suggests bool",
			col:  ddl.Column{RawType: "BOOLEAN"},
			want: "bool",
		},
		{
			name: "VARCHAR(255) suggests string",
			col:  ddl.Column{RawType: "VARCHAR(255)"},
			want: "string",
		},
		{
			name: "unknown type defaults to string",
			col:  ddl.Column{RawType: "JSONB"},
			want: "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := suggestClaimType(tt.col)
			if got != tt.want {
				t.Errorf("suggestClaimType(%q) = %q, want %q", tt.col.RawType, got, tt.want)
			}
		})
	}
}
