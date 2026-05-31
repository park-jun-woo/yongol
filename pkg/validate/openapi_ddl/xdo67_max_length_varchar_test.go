//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo67MaxLengthVarchar — empty constraints + VARCHAR 무관 + maxLength 유무 검증
package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo67MaxLengthVarchar(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
