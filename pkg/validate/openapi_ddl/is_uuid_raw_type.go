//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what isUUIDRawType — DDL 컬럼의 정규화된 raw type이 UUID인지 판별

package openapi_ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// isUUIDRawType reports whether a DDL column's normalised raw type is UUID.
func isUUIDRawType(c ddl.Column) bool {
	t := strings.TrimSpace(c.RawType)
	t = strings.TrimSuffix(t, "[]")
	t = strings.TrimSpace(t)
	return ddl.NormalizePGTypeHead(t) == "UUID"
}
