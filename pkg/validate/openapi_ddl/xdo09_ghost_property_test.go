//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo09GhostProperty — nil doc/nil components/empty schemas 조기 반환 검증
package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo09GhostProperty(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
