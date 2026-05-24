//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what TestIsUUIDRawType — DDL RawType UUID 판별 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsUUIDRawType(t *testing.T) {
	cases := []struct {
		raw  string
		want bool
	}{
		{"UUID", true},
		{"uuid", true},
		{" UUID ", true},
		{"TEXT", false},
		{"BIGINT", false},
		{"VARCHAR(255)", false},
		{"BOOLEAN", false},
	}
	for _, c := range cases {
		got := isUUIDRawType(ddl.Column{RawType: c.raw})
		if got != c.want {
			t.Errorf("isUUIDRawType(%q) = %v, want %v", c.raw, got, c.want)
		}
	}
}
